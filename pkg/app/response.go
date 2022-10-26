package app

import (
	"fujitech/web/web-go/go-challenge/pkg/e"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type ResponseOk struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseNG struct {
	Status bool   `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

// Response setting gin.JSON in OK case
func (g *Gin) ResponseSuccess(data interface{}) {
	g.C.JSON(http.StatusOK, ResponseOk{
		Status: true,
		Data:   data,
	})
}

// Response setting gin.JSON in NG case
func (g *Gin) ResponseError(errCode int, msg interface{}) {
	if reflect.TypeOf(msg) == nil {
		g.C.JSON(http.StatusBadRequest, ResponseNG{
			Status: false,
			Code:   errCode,
			Msg:    e.GetMsg(errCode),
		})
	} else {
		g.C.JSON(http.StatusBadRequest, ResponseNG{
			Status: false,
			Code:   errCode,
			Msg:    msg.(string),
		})
	}
}
