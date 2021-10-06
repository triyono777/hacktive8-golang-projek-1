package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hacktive8-golang-projek-1/database"
	"hacktive8-golang-projek-1/helpers"
	"hacktive8-golang-projek-1/models"
	"net/http"
	"strconv"
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
func GetSingleTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Todo models.Todo
	todoId, _ := strconv.Atoi(c.Param("todoId"))

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

	err := db.Debug().Where("id = ?",todoId).First(&Todo).Error

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
		"message": "Success get single data",
		"data":    Todo,
	})

}
func UpdateTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Todo := models.Todo{}
	todoId, _ := strconv.Atoi(c.Param("todoId"))

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
	Todo.ID = uint(todoId)

	err := db.Debug().Model(&Todo).Where("id = ?", todoId).Updates(models.Todo{
		Title: Todo.Title,
		Done:  Todo.Done,
	}).Error

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
		"message": "Success update data",
		"data":    Todo,
	})

}
func DeleteTodo(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Todo := models.Todo{}
	todoId, _ := strconv.Atoi(c.Param("todoId"))

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
	Todo.ID = uint(todoId)

	err := db.Debug().Model(&Todo).Where("id = ?", todoId).Delete(&Todo).Error

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
		"message": "Success delete data",
	})

}
