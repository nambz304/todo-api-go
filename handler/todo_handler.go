package handler

import (
	"fmt"
	"net/http"
	"todo-api/models"
	"todo-api/service"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo, err := service.AddTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTodo)
}

func GetTodos(c *gin.Context) {
	var todo models.Todo
	newTodo, err := service.GetTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTodo)
}

func DeleteTodoByID(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("c là: ", *c)
	// id, _ := strconv.ParseUint(idStr, 10, 32)
	idStr := todo.ID
	fmt.Println("ID cần xoá: ", idStr)
	// id, _ := strconv.ParseUint(idStr, 10, 32) // Chuyển string → uint
	service.DeleteTodoByID_(uint(idStr))
	// ...
}

func UpdateTodoByID(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo, err := service.UpdateTodoByID_(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newTodo)
}
