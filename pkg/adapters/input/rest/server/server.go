package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/olegshishkin/financier/api"
)

type Server struct {
	router *gin.Engine
}

func NewServer(middleware ...gin.HandlerFunc) *Server {
	router := gin.New()
	router.Use(middleware...)

	return &Server{
		router: router,
	}
}

func (s *Server) RegisterRoutes(api api.ServerInterface) {
	s.configureSwagger()
	accounts := s.router.Group("/accounts")
	{
		accounts.POST("", api.AddAccount)
		accounts.GET("", api.GetAllAccounts)
	}
}

func (s *Server) Start() {
	if err := s.router.Run(); err != nil {
		panic(errors.Wrap(err, "Web Server error"))
	}
}

func (s *Server) configureSwagger() {
	// ginSwagger.WrapHandler(
	// 	swaggerFiles.Handler,
	// 	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	// 	ginSwagger.DefaultModelsExpandDepth(-1))
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("swagger.yaml")))
}
