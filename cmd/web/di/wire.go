//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/olegshishkin/financier/pkg/adapters/input/rest/server"
)

func Wire() *server.Server {
	panic(wire.Build(WebAppProviderSetStub))
}
