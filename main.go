package main

import (
	"hacktive8-golang-projek-1/database"
	"hacktive8-golang-projek-1/router"
)

// @title Todos API
// @version 1.0
// @description This is a documentation service for todo apps
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email yono.tri@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /


// CreateTodo godoc
// @Summary Show the status of server.
// @Description create new todos with input payload.
// @Tags todos
// @Accept json
// @Produce json
// Param todos string true "title type string"  // masih belum sesuai ekpektasi
// @Success 200 {object} map[string]interface{}
// @Router /todos [post]


//getTodo godoc
// @Summary Show the status of server.
// @Description create new todos with input payload.
// @Tags getTodos
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /todos [get]
func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":9000")
}
