package v1

import (
	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/service/transaction_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionForm struct {
	UserId       string
	PartnerId    string `form:"partner_id" valid:"MaxSize(36)"`
	BankId       string `form:"bank_id" valid:"MaxSize(36)"`
	BankBranchId string `form:"bank_branch_id" valid:"MaxSize(36)"`
	Amount       int64  `form:"amount"`
	Fee          int64  `form:"fee"`
	Message      string `form:"message" valid:"MaxSize(100)"`
	Currency     int    `form:"currency"`
}

// @Summary Get all transactions of user
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions [get]
func GetTransactions(c *gin.Context) {
	var appG = app.Gin{C: c}

	// Get current user
	h := c.Request.Header.Get("token")
	currentUser, err := util.ParseToken(h)
	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
		return
	}

	transactionService := transaction_service.Transaction{
		UserId: util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
	}

	appG.ResponseSuccess(transactionService.GetAll())
}

// @Summary Get transaction details
// @Produce  json
// @Param id path string true "Transaction ID"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions/:id [get]
func GetTransaction(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Param("id")

	tService := transaction_service.Transaction{
		ID: id,
	}
	// Check exist
	if exist := tService.CheckExist(); !exist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_NOT_EXIST_TRANSACTION))
		return
	}

	appG.ResponseSuccess(tService.GetId())
}

// @Summary Search transaction
// @Produce  json
// @Param start_date path string true "Start date"
// @Param end_date path string true "End date"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions/search/:start_date/:end_date [get]
func GetTransactionByDate(c *gin.Context) {
	var appG = app.Gin{C: c}

	startDate := c.Param("start_date")
	endDate := c.Param("end_date")

	appG.ResponseSuccess(transaction_service.GetTransactionByDate(startDate, endDate))
}

// @Summary Register new transaction
// @Produce  json
// @Param partner_id path string false "Partner Id"
// @Param bank_id path string false "Bank Id"
// @Param bank_branch_id path string false "Bank branch Id"
// @Param amount path int64 false "Amount"
// @Param fee path int64 false "Fee"
// @Param message path string false "Message"
// @Param currency path int false "Currency"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions [post]
func RegisterTransaction(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TransactionForm
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

	transactionService := transaction_service.Transaction{
		UserId:       util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
		PartnerId:    form.PartnerId,
		BankId:       form.BankId,
		BankBranchId: form.BankBranchId,
		Amount:       form.Amount,
		Fee:          form.Fee,
		Message:      form.Message,
		Currency:     form.Currency,
	}

	// Validate and addition
	if err := transactionService.Add(); err != nil {
		appG.ResponseError(e.ERROR, err)
		return
	}

	appG.ResponseSuccess(e.GetMsg(e.SUCCESS_ADD_TRANSACTION))
}

// @Summary Store edition transaction
// @Produce  json
// @Param id path string false "Transaction Id"
// @Param partner_id path string false "Partner Id"
// @Param bank_id path string false "Bank Id"
// @Param bank_branch_id path string false "Bank branch Id"
// @Param amount path int64 false "Amount"
// @Param fee path int64 false "Fee"
// @Param message path string false "Message"
// @Param currency path int false "Currency"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions/:id [put]
func StoreTransaction(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TransactionForm
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

	tService := transaction_service.Transaction{
		ID:           id,
		PartnerId:    form.PartnerId,
		BankId:       form.BankId,
		BankBranchId: form.BankBranchId,
		Amount:       form.Amount,
		Fee:          form.Fee,
		Message:      form.Message,
		Currency:     form.Currency,
		UserId:       util.AesDecrypt(currentUser.Id, setting.AppSetting.Key),
	}

	// Check exist
	if exist := tService.CheckExist(); !exist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_NOT_EXIST_TRANSACTION))
		return
	}

	// Validate and update
	if err := tService.Edit(); err != nil {
		appG.ResponseError(e.ERROR, err)
		return
	}

	appG.ResponseSuccess(e.GetMsg(e.SUCCESS_EDIT_TRANSACTION))
}

// @Summary Update transaction status
// @Produce  json
// @Param id path string true "Transaction Id"
// @Param status path int true "Status (0-Create; 1-Pending; 2-Success)"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions/status/:id/:status [post]
func StoreStatusTransaction(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Param("id")

	status, err := strconv.Atoi(c.Param("status"))
	if err != nil || (status > 2 && status < 0) {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_NOT_STATUS_TRANSACTION))
	}

	transaction_service.UpdateStatus(id, status)
	appG.ResponseSuccess(e.GetMsg(e.SUCCESS_UPDATE_STATUS_TRANSACTION))
}

// @Summary Delete a transaction
// @Produce  json
// @Param id path string true "Transaction Id"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /transactions/:id [delete]
func DeleteTransaction(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := c.Param("id")

	tService := transaction_service.Transaction{
		ID: id,
	}

	// Check exist
	if exist := tService.CheckExist(); !exist {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_NOT_EXIST_TRANSACTION))
		return
	}

	// Remove
	if err := tService.Delete(); err != nil {
		appG.ResponseError(e.ERROR, err)
		return
	}

	appG.ResponseSuccess(e.GetMsg(e.SUCCESS_DELETE_TRANSACTION))
}
