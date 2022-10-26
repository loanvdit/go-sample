package v1

import (
	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/service/partner_service"

	"github.com/gin-gonic/gin"
)

type PartnerForm struct {
	UserId      string
	Name        string `form:"name" valid:"Required;MaxSize(100)"`
	BankAccount string `form:"bank_account" valid:"Required;MaxSize(30)"`
	Address     string `form:"address" valid:"MaxSize(300)"`
	Email       string `form:"email" valid:"MaxSize(30)"`
	Phone       string `form:"phone" valid:"MaxSize(20)"`
}

// @Summary Get all partner of user
// @Produce  json
// @Param token path string true "Token"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /partners [get]
func GetPartners(c *gin.Context) {
	var appG = app.Gin{C: c}

	// Get current user
	h := c.Request.Header.Get("token")
	currentUser, err := util.ParseToken(h)
	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
		return
	}

	partnerService := partner_service.Partner{
		UserId: util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
	}

	appG.ResponseSuccess(partnerService.GetAll())
}

// @Summary Register new partner
// @Produce  json
// @Param name query string true "Name"
// @Param bank_account query string true "Bank account"
// @Param address query string true "Address"
// @Param email query string true "Email"
// @Param phone query string true "Phone"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /partners [post]
func RegisterPartner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PartnerForm
	)

	// Get current user
	h := c.Request.Header.Get("token")
	currentUser, err := util.ParseToken(h)
	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
		return
	}

	// Validate form
	errCode, Msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.ResponseError(errCode, Msg)
		return
	}

	partnerService := partner_service.Partner{
		Name:        form.Name,
		BankAccount: form.BankAccount,
		Address:     form.Address,
		Email:       form.Email,
		Phone:       form.Phone,
		UserId:      util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
	}

	// Validate and add
	if exist := partnerService.CheckValid(); exist {
		if err := partnerService.Add(); err != nil {
			appG.ResponseError(e.ERROR, err)
			return
		}
	} else {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_PARTNER_EXIST))
		return
	}

	appG.ResponseSuccess(partnerService.GetCurrentPartner())
}

// @Summary Store edition partner
// @Produce  json
// @Param name query string true "Name"
// @Param bank_account query string true "Bank account"
// @Param address query string true "Address"
// @Param email query string true "Email"
// @Param phone query string true "Phone"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /partners [put]
func StorePartner(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form PartnerForm
	)
	id := c.Param("id")

	// Get current user
	h := c.Request.Header.Get("token")
	currentUser, err := util.ParseToken(h)
	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
		return
	}

	// Validate form
	errCode, Msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.ResponseError(errCode, Msg)
		return
	}

	partnerService := partner_service.Partner{
		Id:          id,
		Name:        form.Name,
		BankAccount: form.BankAccount,
		Address:     form.Address,
		Email:       form.Email,
		Phone:       form.Phone,
		UserId:      util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
	}

	// Check exist
	if exist := partnerService.CheckExist(); !exist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_PARTNER_NOT_EXIST))
		return
	}

	// Validate and update
	if err := partnerService.Edit(); err != nil {
		appG.ResponseError(e.ERROR, err)
		return
	}

	appG.ResponseSuccess(partnerService.GetCurrentPartner())
}

// @Summary Delete a partner
// @Produce  json
// @Param id path string true "id"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /partners/:id [delete]
func DeletePartner(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Param("id")

	partnerService := partner_service.Partner{
		Id: id,
	}

	// Check exist
	if exist := partnerService.CheckExist(); !exist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_PARTNER_NOT_EXIST))
		return
	}

	// Remove
	if err := partnerService.Delete(); err != nil {
		appG.ResponseError(e.ERROR, err)
		return
	}

	appG.ResponseSuccess(e.GetMsg(e.SUCCESS_DELETE_USER))
}
