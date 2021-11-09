package controllers

import (
	"fmt"
	"hacktive8-golang-projek-1/database"
	_ "hacktive8-golang-projek-1/docs"
	"hacktive8-golang-projek-1/helpers"
	"hacktive8-golang-projek-1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/http-swagger"
)

var (
	appJSON = "application/json"
)

// CreateTodo godoc
// @Summary Create todos.
// @Description create new todos with input payload.
// @Tags todos
// @Accept json
// @Produce json
// @Param   Title     query    string     true        "title todo"
// @Success 200 {object} map[string]interface{}
// @Router /todos [post]
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

// GetTodos getTodo godoc
//@Summary get todos
//@Description create new todos with input payload.
//@Tags todos
//@Accept json
//@Produce json
//@Success 200 {object} map[string]interface{}
//@Router /todos [get]
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
	if len(Todos) == 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Data kosong",
			"data":    Todos,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Success get data",
			"data":    Todos,
		})
	}

}

// GetSingleTodo  get single Todo godoc
//@Summary get single todos.
//@Description get single todo.
//@Tags todos
//@Accept json
//@Produce json
//@Param   todoId     path    string     true        "todo id"
//@Success 200 {object} map[string]interface{}
//@Router /todos/{todoId} [get]
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

	err := db.Debug().Where("id = ?", todoId).First(&Todo).Error

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

// UpdateTodo godoc
// @Summary Update todos.
// @Description update new todos with input payload.
// @Tags todos
// @Accept json
// @Produce json
// @Param   todoId     path    string     true        "todo id"
// @Param   Title     query    string     true        "title todo"
// @Param   Done     query    bool     true        "todo done or yet "
// @Success 200 {object} map[string]interface{}
// @Router /todos/{todoId} [put]
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

// DeleteTodo godoc
// @Summary Delete todo.
// @Description delete todo.
// @Tags todos
// @Accept json
// @Produce json
// @Param   todoId     path    string     true        "todo id"
// @Success 200 {object} map[string]interface{}
// @Router /todos/{todoId} [delete]
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
