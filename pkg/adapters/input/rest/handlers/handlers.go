package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/api"
)

type HandlerDelegate struct {
	accountHdl AccountHTTPRequestHandler
}

func NewHandlerDelegate(accountHdl AccountHTTPRequestHandler) *HandlerDelegate {
	return &HandlerDelegate{
		accountHdl: accountHdl,
	}
}

func (h *HandlerDelegate) AddAccount(ctx *gin.Context) {
	h.accountHdl.AddAccount(ctx)
}

func (h *HandlerDelegate) GetAllAccounts(ctx *gin.Context) {
	h.accountHdl.GetAllAccounts(ctx)
}

//nolint:revive,stylecheck
func (h *HandlerDelegate) FindAccountById(ctx *gin.Context, id api.Id) {
	h.accountHdl.FindAccountByID(ctx, id)
}
