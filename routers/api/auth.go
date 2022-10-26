package api

import (
	"crypto/tls"
	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/service/auth_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

type auth struct {
	Email    string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

type key struct {
	ActiveKey string `valid:"Required; MaxSize(5)"`
}

// @Summary Create key to login
// @Produce  json
// @Param email query string true "email"
// @Param password query string true "password"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	email := c.PostForm("email")
	password := c.PostForm("password")

	a := auth{Email: email, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.ResponseError(e.ERROR_AUTH_INVALID_PARAM, valid.Errors[0].Message)
		return
	}

	authService := auth_service.Auth{Email: email, Password: password}
	isExist, id, verifyKey, err := authService.Check()
	if err != nil || !isExist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_ACCOUNT_NOT_EXIST))
		return
	}

	// Sending verify key to email
	go SendVerifyKey(email, verifyKey)

	// Return user id
	appG.ResponseSuccess(map[string]string{
		"user_id": id,
	})
}

// @Summary Create user token
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /auth/:id [post]
func GetAuthVerify(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := c.Param("id")

	activekey := c.PostForm("verify_key")

	a := key{ActiveKey: activekey}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.ResponseError(e.INVALID_PARAMS, valid.Errors[0].Message)
		return
	}

	authService := auth_service.Auth{Id: id, VerifyKey: activekey}
	isExist, err := authService.ValidKey()

	if err != nil {
		appG.ResponseError(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil)
		return
	}

	if !isExist {
		appG.ResponseError(e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(id, authService.Email, authService.Password)
	if err != nil {
		appG.ResponseError(e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.ResponseSuccess(map[string]string{
		"token": token,
	})
}

// Send serial number
func SendVerifyKey(email string, serial string) {
	m := gomail.NewMessage()

	m.SetHeader("From", setting.SmtpSetting.User)

	m.SetHeader("To", email)

	m.SetHeader("Subject", "Verify serial number email")

	m.SetBody("text/plain", serial)

	d := gomail.NewDialer(setting.SmtpSetting.Host, setting.SmtpSetting.Port, setting.SmtpSetting.User, setting.SmtpSetting.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
