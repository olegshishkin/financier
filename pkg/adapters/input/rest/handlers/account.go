package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/olsh-go-utils/types"

	"github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/mapper"
	"github.com/olegshishkin/financier/pkg/core/ports/input"
)

type AccountHTTPRequestHandler interface {
	GetAllAccounts(c *gin.Context)
	AddAccount(c *gin.Context)
	FindAccountByID(c *gin.Context, id v1.ID)
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
	var rq v1.NewAccountInput

	if err := ctx.ShouldBindJSON(&rq); err != nil {
		rest.Err(ctx, http.StatusBadRequest, rest.Tech, err)

		return
	}

	acc, err := h.accountSvc.Create(rq.Name, types.PointerVal(rq.Comment))
	if err != nil {
		rest.Err(ctx, http.StatusBadRequest, rest.Business, err)

		return
	}

	ctx.JSON(http.StatusCreated, mapper.AccountToAccountOut(acc))
}

func (h *AccountHandler) GetAllAccounts(ctx *gin.Context) {
	accounts, err := h.accountSvc.GetAll()
	if err != nil {
		rest.Err(ctx, http.StatusInternalServerError, rest.Tech, err)

		return
	}

	ctx.JSON(http.StatusOK, mapper.AccountsToAccountsOut(accounts))
}

func (h *AccountHandler) FindAccountByID(ctx *gin.Context, _ v1.ID) {
	ctx.JSON(http.StatusNotImplemented, "")
}
