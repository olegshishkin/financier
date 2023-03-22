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
	WebAppProviderSetStub = wire.NewSet(
		provideServer,
		provideServerMiddlewares,
		provideHandlerDelegate,
		provideSwaggerHandler,
		accountProviderSetStub,
		provideWebLogger,
		provideSourceLogger,
		wire.Bind(new(handlers.SwaggerHTTPRequestHandler), new(*handlers.SwaggerHandler)),
		wire.Bind(new(handlers.AccountHTTPRequestHandler), new(*handlers.AccountHandler)),
		wire.Bind(new(logger.Logger), new(*zerolog.Wrapper)),
		wire.Bind(new(v1.ServerInterface), new(*handlers.HandlerDelegate)),
	)

	//nolint:gochecknoglobals
	accountProviderSetStub = wire.NewSet(
		provideAccountHandler,
		provideAccountService,
		provideAccountStorageStub,
		wire.Bind(new(input.AccountService), new(*services.AccountService)),
		wire.Bind(new(output.AccountStorage), new(*stub.AccountStorageStub)),
	)
)

func provideAccountStorageStub() *stub.AccountStorageStub {
	var stg *stub.AccountStorageStub

	accStorageOnce.Do(func() {
		stg = stub.NewAccountStorageStub()
	})

	return stg
}
