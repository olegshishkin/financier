package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/api/v1"
)

type HandlerDelegate struct {
	swaggerHdl SwaggerHTTPRequestHandler
	accountHdl AccountHTTPRequestHandler
}

func NewHandlerDelegate(swaggerHdl SwaggerHTTPRequestHandler, accountHdl AccountHTTPRequestHandler) *HandlerDelegate {
	return &HandlerDelegate{
		swaggerHdl: swaggerHdl,
		accountHdl: accountHdl,
	}
}

func (h *HandlerDelegate) GetSwagger(ctx *gin.Context) {
	h.swaggerHdl.GetSwagger(ctx)
}

func (h *HandlerDelegate) AddAccount(ctx *gin.Context) {
	h.accountHdl.AddAccount(ctx)
}

func (h *HandlerDelegate) GetAllAccounts(ctx *gin.Context) {
	h.accountHdl.GetAllAccounts(ctx)
}

func (h *HandlerDelegate) FindAccountByID(ctx *gin.Context, id v1.ID) {
	h.accountHdl.FindAccountByID(ctx, id)
}
