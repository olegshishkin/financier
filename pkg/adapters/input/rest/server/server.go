package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olegshishkin/go-logger"

	v1 "github.com/olegshishkin/financier/api/v1"
	"github.com/olegshishkin/financier/config"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest"
	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
)

type Server struct {
	router *gin.Engine
	log    logger.Logger
	host   string
	port   string
}

func NewServer(cfg *config.Config, log logger.Logger) *Server {
	if mode := cfg.Server.Mode; mode != "" {
		gin.SetMode(mode)
	}

	return &Server{
		router: gin.New(),
		log:    log,
		host:   cfg.Server.Host,
		port:   cfg.Server.Port,
	}
}

func (s *Server) RegisterHandlers(handlers *handlers.HandlerDelegate, mdl *Middlewares) {
	s.router.Use(mdl.Recovery, mdl.Logging, mdl.OpenAPIValidation)

	opts := v1.GinServerOptions{
		BaseURL:     "",
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
	var addr []string

	if port := s.port; port != "" {
		addr = append(addr, fmt.Sprintf("%v:%v", s.host, port))
	}

	if err := s.router.Run(addr...); err != nil {
		s.log.Fatal(err, "Web Server error")
	}
}
