package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hacktive8-golang-projek-1/models"
	"log"
)

var (
	host     = "localhost"
	user     = "root"
	password = "Triyono7"
	dbPort   = "3306"
	dbName   = "db_todos"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbName)
	dsn := config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting", err)
	}
	fmt.Println("sukses koneksi ke db")
	err = db.Debug().AutoMigrate(models.Todo{})
	if err != nil {
		return
	}
}
func GetDB() *gorm.DB {
	return db
}
