package router

import "github.com/gin-gonic/gin"

func StartApp() *gin.Engine {
	todoRouter := gin.Default()

	{
		todoRouter.POST("/todos")
	}

	return todoRouter
}
