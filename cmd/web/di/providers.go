package di

import (
	"sync"

	"github.com/olegshishkin/go-logger"
	logzero "github.com/olegshishkin/go-logger-zerolog"
	logzeroexample "github.com/olegshishkin/go-logger-zerolog/example"
	"github.com/rs/zerolog"

	"github.com/olegshishkin/financier/config"
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
	swgHandlerOnce sync.Once

	//nolint:gochecknoglobals
	accHandlerOnce sync.Once

	//nolint:gochecknoglobals
	accServiceOnce sync.Once

	//nolint:gochecknoglobals
	accStorageOnce sync.Once

	//nolint:gochecknoglobals
	middlewareOnce sync.Once

	//nolint:gochecknoglobals
	webLogOnce sync.Once

	//nolint:gochecknoglobals
	sourceLogOnce sync.Once

	//nolint:gochecknoglobals
	configOnce sync.Once
)

func provideServer(
	cfg *config.Config,
	log logger.Logger,
	handlers *handlers.HandlerDelegate,
	mdl *server.Middlewares,
) *server.Server {
	var srv *server.Server

	serverOnce.Do(func() {
		srv = server.NewServer(cfg, log)
		srv.RegisterSwaggerHandler(handlers, mdl)
		srv.RegisterHandlers(handlers, mdl)
	})

	return srv
}

func provideServerMiddlewares(log logger.Logger) *server.Middlewares {
	var mdl *server.Middlewares

	middlewareOnce.Do(func() {
		mdl = server.NewMiddlewares(log)
	})

	return mdl
}

func provideHandlerDelegate(
	sh handlers.SwaggerHTTPRequestHandler,
	ah handlers.AccountHTTPRequestHandler,
) *handlers.HandlerDelegate {
	var hd *handlers.HandlerDelegate

	hdlDelegateOnce.Do(func() {
		hd = handlers.NewHandlerDelegate(sh, ah)
	})

	return hd
}

func provideSwaggerHandler(log logger.Logger) *handlers.SwaggerHandler {
	var hdl *handlers.SwaggerHandler

	swgHandlerOnce.Do(func() {
		hdl = handlers.NewSwaggerHandler(log)
	})

	return hdl
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

func provideWebLogger(zeroLogger *zerolog.Logger) *logzero.Wrapper {
	var log *logzero.Wrapper

	webLogOnce.Do(func() {
		log = logzero.From(logzeroexample.Web(zeroLogger))
	})

	return log
}

func provideSourceLogger(cfg *config.Config) *zerolog.Logger {
	var log *zerolog.Logger

	sourceLogOnce.Do(func() {
		writer, err := logzero.NewLogWriterBuilder().
			WithConsole(logzeroexample.ConsoleWriter()).
			Build()
		if err != nil {
			panic(err)
		}

		lvl, err := logger.ParseLevel(cfg.Logging.Level)
		if err != nil {
			panic(err)
		}

		log = logzeroexample.Base(writer, lvl)
	})

	return log
}

func provideConfig() *config.Config {
	var cfg *config.Config

	configOnce.Do(func() {
		cfg = config.ReadConfig()
	})

	return cfg
}
