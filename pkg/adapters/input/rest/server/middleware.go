package server

import (
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"
	loggin "github.com/olegshishkin/go-logger-gin"

	v1 "github.com/olegshishkin/financier/api/v1"
)

type Middlewares struct {
	Recovery          gin.HandlerFunc
	Logging           gin.HandlerFunc
	OpenAPIValidation gin.HandlerFunc
}

func NewMiddlewares(log logger.Logger) *Middlewares {
	return &Middlewares{
		Recovery:          gin.Recovery(),
		Logging:           loggin.WebServerLogger(log),
		OpenAPIValidation: validationMiddleware(log),
	}
}

func validationMiddleware(log logger.Logger) gin.HandlerFunc {
	swagger, err := v1.GetSwagger()
	if err != nil {
		log.Fatal(err, "OpenAPI schema validator hasn't been created")
	}

	return middleware.OapiRequestValidator(swagger)
}
