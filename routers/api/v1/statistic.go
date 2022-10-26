package v1

import (
	"fmt"
	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/service/statistic_service"

	"github.com/gin-gonic/gin"
)

// @Summary Get details report
// @Produce  json
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /statistic/:id [get]
func GetStatisticReport(c *gin.Context) {
	var appG = app.Gin{C: c}

	// Get current user
	h := c.Request.Header.Get("token")
	currentUser, err := util.ParseToken(h)
	if err != nil {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_TOKEN_TIMEOUT))
		return
	}

	result, total := statistic_service.Get(util.AesDecrypt(currentUser.Id, setting.AppSetting.Key))

	if total == 0 {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_STATISTIC_NO_RESULT))
		return
	}

	fmt.Printf("%+v\n", result)
	appG.ResponseSuccess(result)
}
