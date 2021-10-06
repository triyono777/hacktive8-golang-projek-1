package router

import (
	"github.com/gin-gonic/gin"
	"hacktive8-golang-projek-1/controllers"
)

func StartApp() *gin.Engine {
	todoRouter := gin.Default()

	{
		todoRouter.POST("/todos",controllers.CreateTodo)
		todoRouter.GET("/todos",controllers.GetTodos)
	}

	return todoRouter
}
