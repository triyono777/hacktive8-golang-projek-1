package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hacktive8-golang-projek-1/database"
	"hacktive8-golang-projek-1/helpers"
	"hacktive8-golang-projek-1/models"
	"net/http"
)

var (
	appJSON = "application/json"
)

func CreateTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Todo := models.Todo{}

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Todo)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Todo)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	err := db.Debug().Create(&Todo).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Todo,
	})

}
func GetTodos(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Todos []models.Todo

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Todos)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Todos)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	err := db.Debug().Find(&Todos).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Todos,
	})

}
