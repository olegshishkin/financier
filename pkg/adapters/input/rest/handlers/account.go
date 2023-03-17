package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/api"
	"github.com/olegshishkin/financier/pkg/core/ports/input"
)

type AccountHTTPRequestHandler interface {
	GetAllAccounts(c *gin.Context)
	AddAccount(c *gin.Context)
	FindAccountByID(c *gin.Context, id api.Id)
}

type AccountHandler struct {
	accountSvc input.AccountService
}

func NewAccountHandler(as input.AccountService) *AccountHandler {
	return &AccountHandler{
		accountSvc: as,
	}
}

func (h *AccountHandler) AddAccount(ctx *gin.Context) {
	var rq api.NewAccount

	if err := ctx.ShouldBindJSON(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	acc, err := h.accountSvc.Create(rq.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, acc)
}

func (h *AccountHandler) GetAllAccounts(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "")
}

func (h *AccountHandler) FindAccountByID(ctx *gin.Context, id api.Id) {
	ctx.JSON(http.StatusNotImplemented, "")
}
