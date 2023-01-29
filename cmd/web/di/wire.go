//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func WireStubs() *gin.Engine {
	panic(wire.Build(WebAppProviderSetStub))
}
