package di

import (
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/routes"
	"github.com/olegshishkin/financier/pkg/core/ports/input"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/services"
)

var (
	//nolint:gochecknoglobals
	ginRouterOnce sync.Once

	//nolint:gochecknoglobals
	accHandlerOnce sync.Once

	//nolint:gochecknoglobals
	accServiceOnce sync.Once

	//nolint:gochecknoglobals
	accStorageOnce sync.Once
)

func provideGinRouter(accHdl *handlers.AccountHandler) *gin.Engine {
	var router *gin.Engine

	ginRouterOnce.Do(func() {
		recHdl := gin.Recovery()
		router = routes.GinRouter(accHdl, recHdl)
	})

	return router
}

func provideAccountHandler(as input.AccountService) *handlers.AccountHandler {
	var hdl *handlers.AccountHandler

	accHandlerOnce.Do(func() {
		hdl = handlers.NewAccountHandler(as)
	})

	return hdl
}

func provideAccountService(as output.AccountStorage) *services.AccountService {
	var svc *services.AccountService

	accServiceOnce.Do(func() {
		svc = services.NewAccountService(as)
	})

	return svc
}
