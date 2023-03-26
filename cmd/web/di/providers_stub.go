package di

import (
	"github.com/google/wire"
	"github.com/olegshishkin/go-logger"
	zerolog "github.com/olegshishkin/go-logger-zerolog"

	"github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
	"github.com/olegshishkin/financier/pkg/adapters/output/persistence/stub"
	"github.com/olegshishkin/financier/pkg/core/ports/input"
	"github.com/olegshishkin/financier/pkg/core/ports/output"
	"github.com/olegshishkin/financier/pkg/core/services"
)

var (
	//nolint:gochecknoglobals
	loggerProviderSetStub = wire.NewSet(
		provideWebLogger,
		provideSourceLogger,
		wire.Bind(new(logger.Logger), new(*zerolog.Wrapper)),
	)

	//nolint:gochecknoglobals
	accountProviderSetStub = wire.NewSet(
		provideAccountStorageStub,
		provideAccountService,
		provideAccountHandler,
		wire.Bind(new(input.AccountService), new(*services.AccountService)),
		wire.Bind(new(output.AccountStorage), new(*stub.AccountStorageStub)),
	)

	//nolint:gochecknoglobals
	handlersProviderSetStub = wire.NewSet(
		accountProviderSetStub,
		provideSwaggerHandler,
		provideHandlerDelegate,
		wire.Bind(new(handlers.SwaggerHTTPRequestHandler), new(*handlers.SwaggerHandler)),
		wire.Bind(new(handlers.AccountHTTPRequestHandler), new(*handlers.AccountHandler)),
	)

	//nolint:gochecknoglobals
	WebAppProviderSetStub = wire.NewSet(
		provideConfig,
		loggerProviderSetStub,
		handlersProviderSetStub,
		provideServerMiddlewares,
		provideServer,
		wire.Bind(new(v1.ServerInterface), new(*handlers.HandlerDelegate)),
	)
)

func provideAccountStorageStub() *stub.AccountStorageStub {
	var stg *stub.AccountStorageStub

	accStorageOnce.Do(func() {
		stg = stub.NewAccountStorageStub()
	})

	return stg
}
