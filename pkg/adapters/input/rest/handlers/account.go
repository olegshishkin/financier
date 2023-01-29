package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/pkg/core/ports/input"
)

type AccountHandler struct {
	accountSvc input.AccountService
}

type createAccountRq struct {
	name string
}

func NewAccountHandler(as input.AccountService) *AccountHandler {
	return &AccountHandler{
		accountSvc: as,
	}
}

func (h *AccountHandler) CreateAccount(ctx *gin.Context) {
	var rq createAccountRq

	if err := ctx.ShouldBindJSON(&rq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	acc, err := h.accountSvc.Create(rq.name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, acc)
}
