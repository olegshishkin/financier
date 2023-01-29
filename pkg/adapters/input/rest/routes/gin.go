package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/olegshishkin/financier/pkg/adapters/input/rest/handlers"
)

func GinRouter(accHdl *handlers.AccountHandler, middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middleware...)

	accountsRouter := router.Group("/accounts")
	{
		accountsRouter.POST("/create", accHdl.CreateAccount)
	}

	return router
}
