package route

import (
	"config/handler"
	"github.com/gin-gonic/gin"
)

func GetHttpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handler.Logger())
	user := router.Group("/api/v1/user")
	{
		// add user
		user.POST("/add", handler.AddUserHandler)
	}

	return router
}
