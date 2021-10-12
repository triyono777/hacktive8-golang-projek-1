package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"hacktive8-golang-projek-1/controllers"
	_ "hacktive8-golang-projek-1/docs"
)


func StartApp() *gin.Engine {
	todoRouter := gin.Default()
	url:= ginSwagger.URL("http://localhost:9000/docs/doc.json")

	{
		todoRouter.GET("/docs/*any",ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		todoRouter.POST("/todos",controllers.CreateTodo)
		todoRouter.GET("/todos",controllers.GetTodos)
		todoRouter.GET("/todos/:todoId",controllers.GetSingleTodo)
		todoRouter.PUT("/todos/:todoId",controllers.UpdateTodo)
		todoRouter.DELETE("/todos/:todoId",controllers.DeleteTodo)

	}


	return todoRouter
}
