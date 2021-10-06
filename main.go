package main

import (
	"hacktive8-golang-projek-1/database"
	"hacktive8-golang-projek-1/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":9000")
}
