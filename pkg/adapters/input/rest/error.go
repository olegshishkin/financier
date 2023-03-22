package rest

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/olegshishkin/financier/api/v1"
)

const (
	_ ErrCode = iota
	Tech
	Business
)

type ErrCode int8

func Err(ctx *gin.Context, statusCode int, code ErrCode, err error) {
	e := &v1.Error{
		Code:    int8(code),
		Message: err.Error(),
	}
	ctx.JSON(statusCode, e)
}
