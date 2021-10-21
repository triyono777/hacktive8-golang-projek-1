package main

import (
	_ "github.com/swaggo/http-swagger"
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
// @host localhost:9001
// @BasePath /

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":9001")
}
