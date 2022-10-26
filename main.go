package main

import (
	"fmt"
	"log"
	"net/http"

	"fujitech/web/web-go/go-challenge/models"
	"fujitech/web/web-go/go-challenge/pkg/gredis"
	"fujitech/web/web-go/go-challenge/pkg/logging"
	"fujitech/web/web-go/go-challenge/pkg/setting"
	"fujitech/web/web-go/go-challenge/pkg/util"
	"fujitech/web/web-go/go-challenge/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
