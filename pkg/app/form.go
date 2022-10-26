package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"fujitech/web/web-go/go-challenge/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, interface{}) {
	err := c.Bind(form)
	if err != nil {
		return e.INVALID_PARAMS, nil
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return e.ERROR, nil
	}
	if !check {
		MarkErrors(valid.Errors)
		return e.INVALID_PARAMS, valid.Errors[0].Message
	}

	return e.SUCCESS, nil
}
