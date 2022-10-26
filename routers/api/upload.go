package api

import (
	"github.com/gin-gonic/gin"

	"fujitech/web/web-go/go-challenge/pkg/app"
	"fujitech/web/web-go/go-challenge/pkg/e"
	"fujitech/web/web-go/go-challenge/pkg/logging"
	"fujitech/web/web-go/go-challenge/pkg/upload"
)

// @Summary Import Image
// @Produce  json
// @Param image path string true "image"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} app.ResponseOk
// @Failure 500 {object} app.ResponseNG
// @Router /upload [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.ResponseError(e.ERROR, nil)
		return
	}

	if image == nil {
		appG.ResponseError(e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_UPLOAD_FILE))
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_UPLOAD_FILE))
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.ResponseError(e.ERROR, e.GetMsg(e.ERROR_UPLOAD_FILE))
		return
	}

	appG.ResponseSuccess(map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
