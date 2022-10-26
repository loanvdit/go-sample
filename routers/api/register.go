package api

import (
	"crypto/tls"
	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/service/account_service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

type AccountForm struct {
	Email     string `form:"email" valid:"Required;MaxSize(100)"`
	Password  string `form:"password" valid:"Required;MaxSize(100)"`
	FirstName string `form:"first_name" valid:"Required;MaxSize(30)"`
	LastName  string `form:"last_name" valid:"Required;MaxSize(30)"`
	Phone     string `form:"phone" valid:"Required;MaxSize(20)"`
	Address   string `form:"address" valid:"Required;MaxSize(130)"`
	Avatar    string `form:"avatar"`
}

// @Summary Register new account
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /account/register [post]
func RegisterAccount(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AccountForm
	)

	errCode, Msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.ResponseError(errCode, Msg)
		return
	}

	accountService := account_service.Account{
		Email:     form.Email,
		Password:  form.Password,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Phone:     form.Phone,
		Address:   form.Address,
		Avatar:    form.Avatar,
	}

	if exist := accountService.CheckValid(); exist {
		if err := accountService.Add(); err != nil {
			appG.ResponseError(e.ERROR, err)
			return
		}
		// Send smtp email
		go SendEmailVerify(form.Email)
	} else {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_ACCOUNT_EXIST))
		return
	}

	appG.ResponseSuccess(form)
}

// @Summary Active new user
// @Produce  json
// @Param token path string true "Token"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /account/verify/:token [get]
func VerifyAccount(c *gin.Context) {
	appG := app.Gin{C: c}
	token := util.AesDecrypt(c.Param("token"), setting.AppSetting.Key)

	s := strings.Split(token, "#")

	email := s[0]
	createdAt, err := time.Parse(time.Stamp, s[1])

	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_PARSE_DATE))
		return
	}

	if time.Now().Add(-time.Hour * 1).UTC().After(createdAt) {
		accountService := account_service.Account{Email: email}
		if b := accountService.ActiveAccount(); !b {
			appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_ACTIVE_USER))
			return
		}

		appG.ResponseSuccess(e.GetMsg(e.SUCCESS_ACTIVE_USER))
		return
	}

	appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
}

// Send number verify to email of user
func SendEmailVerify(email string) {
	m := gomail.NewMessage()

	m.SetHeader("From", setting.SmtpSetting.User)

	m.SetHeader("To", email)

	m.SetHeader("Subject", "Verify account email")

	m.SetBody("text/plain", BuildVerifyUrl(email))

	d := gomail.NewDialer(setting.SmtpSetting.Host, setting.SmtpSetting.Port, setting.SmtpSetting.User, setting.SmtpSetting.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// Url active user
func BuildVerifyUrl(email string) string {
	token := util.AesEncrypt(email+"#"+time.Now().UTC().Format(time.Stamp), setting.AppSetting.Key)
	return setting.AppSetting.PrefixUrl + "/account/verify/" + token
}
