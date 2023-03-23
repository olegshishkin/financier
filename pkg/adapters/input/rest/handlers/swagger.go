package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"

	v1 "github.com/olegshishkin/financier/api/v1"
)

type SwaggerHTTPRequestHandler interface {
	GetSwagger(c *gin.Context)
}

type SwaggerHandler struct {
	log     logger.Logger
	handler gin.HandlerFunc
}

func NewSwaggerHandler(log logger.Logger) *SwaggerHandler {
	return &SwaggerHandler{
		log:     log,
		handler: configureSwagger(log),
	}
}

func (h *SwaggerHandler) GetSwagger(ctx *gin.Context) {
	h.handler(ctx)
}

func configureSwagger(log logger.Logger) gin.HandlerFunc {
	swagger, err := v1.GetSwagger()
	if err != nil {
		log.Fatal(err, "OpenAPI specification extracting failed")
	}

	json, err := swagger.MarshalJSON()
	if err != nil {
		log.Fatal(err, "OpenAPI specification hasn't been fetched from the binary Swagger object")
	}

	spec := &swag.Spec{
		Version:          "",
		Host:             "",
		BasePath:         "/",
		Schemes:          []string{},
		Title:            "",
		Description:      "",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(json),
	}

	swag.New()
	swag.Register(spec.InstanceName(), spec)

	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
