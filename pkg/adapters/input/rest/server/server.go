package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"

	v1 "github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
)

const apiRootPath = "/api/v1"

type Server struct {
	router *gin.Engine
	log    logger.Logger
}

func NewServer(log logger.Logger) *Server {
	return &Server{
		router: gin.New(),
		log:    log,
	}
}

func (s *Server) RegisterHandlers(handlers *handlers.HandlerDelegate, mdl *Middlewares) {
	s.router.Use(mdl.Recovery, mdl.Logging, mdl.OpenAPIValidation)

	opts := v1.GinServerOptions{
		BaseURL:     apiRootPath,
		Middlewares: nil,
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			rest.Err(c, statusCode, rest.Tech, err)
		},
	}
	v1.RegisterHandlersWithOptions(s.router, handlers, opts)
}

func (s *Server) RegisterSwaggerHandler(handlers *handlers.HandlerDelegate, mdl *Middlewares) {
	group := s.router.Group("/swagger").Use(mdl.Recovery, mdl.Logging)
	group.GET("/*any", handlers.GetSwagger)
	s.log.Info("Swagger handler has been registered")
}

func (s *Server) Start() {
	if err := s.router.Run(); err != nil {
		s.log.Fatal(err, "Web Server error")
	}
}
