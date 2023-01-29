package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/pkg/core/ports/input"
)

type AccountHandler struct {
	accountSvc input.AccountService
}

func NewAccountHandler(as input.AccountService) *AccountHandler {
	return &AccountHandler{
		accountSvc: as,
	}
}

func (h *AccountHandler) CreateAccount(ctx *gin.Context) {
	acc, err := h.accountSvc.Create(ctx.Param("name"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, acc)
}
