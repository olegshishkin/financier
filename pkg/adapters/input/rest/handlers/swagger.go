package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
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
	content, err := os.ReadFile("api/v1/openapi.yaml")
	if err != nil {
		log.Fatal(err, "OpenAPI specification importing error")
	}

	spec := &swag.Spec{
		Version:          "",
		Host:             "",
		BasePath:         "/",
		Schemes:          []string{},
		Title:            "",
		Description:      "",
		InfoInstanceName: "swagger",
		SwaggerTemplate:  string(content),
	}

	swag.New()
	swag.Register(spec.InstanceName(), spec)

	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
