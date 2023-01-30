package di

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"
	loggin "github.com/olegshishkin/go-logger-gin"
	logzero "github.com/olegshishkin/go-logger-zerolog"
	logzeroexample "github.com/olegshishkin/go-logger-zerolog/example"
	"github.com/rs/zerolog"

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

	//nolint:gochecknoglobals
	logOnce sync.Once
)

func provideGinRouter(accHdl *handlers.AccountHandler, log *zerolog.Logger) *gin.Engine {
	var router *gin.Engine

	ginRouterOnce.Do(func() {
		recHdl := gin.Recovery()
		logHdl := loggin.WebServerLogger(logzero.From(logzeroexample.Web(log)))
		router = routes.GinRouter(accHdl, recHdl, logHdl)
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

func provideLogger() *zerolog.Logger {
	var log *zerolog.Logger

	logOnce.Do(func() {
		writer, err := logzero.NewLogWriterBuilder().
			WithConsole(logzeroexample.ConsoleWriter()).
			Build()
		if err != nil {
			panic(err)
		}

		log = logzeroexample.Base(writer, logger.LogLevel())
	})

	return log
}
