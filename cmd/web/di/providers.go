package di

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"
	loggin "github.com/olegshishkin/go-logger-gin"
	logzero "github.com/olegshishkin/go-logger-zerolog"
	logzeroexample "github.com/olegshishkin/go-logger-zerolog/example"
	"github.com/rs/zerolog"

	"github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/server"
	"github.com/olegshishkin/financier/pkg/core/ports/input"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/services"
)

var (
	//nolint:gochecknoglobals
	serverOnce sync.Once

	//nolint:gochecknoglobals
	hdlDelegateOnce sync.Once

	//nolint:gochecknoglobals
	accHandlerOnce sync.Once

	//nolint:gochecknoglobals
	accServiceOnce sync.Once

	//nolint:gochecknoglobals
	accStorageOnce sync.Once

	//nolint:gochecknoglobals
	logOnce sync.Once
)

func provideServer(apiHandler api.ServerInterface, log *zerolog.Logger) *server.Server {
	var srv *server.Server

	serverOnce.Do(func() {
		recHdl := gin.Recovery()
		logHdl := loggin.WebServerLogger(logzero.From(logzeroexample.Web(log)))
		srv = server.NewServer(recHdl, logHdl)
		srv.RegisterRoutes(apiHandler)
	})

	return srv
}

func provideHandlerDelegate(ah handlers.AccountHTTPRequestHandler) *handlers.HandlerDelegate {
	var hd *handlers.HandlerDelegate

	hdlDelegateOnce.Do(func() {
		hd = handlers.NewHandlerDelegate(ah)
	})

	return hd
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
