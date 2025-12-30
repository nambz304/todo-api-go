package main

import (
	"todo-api/handler"
	"todo-api/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.ConnectDB() // Kết nối DB

	r := gin.Default()
	r.POST("/todos", handler.CreateTodo)
	// Thêm route khác...
	r.GET("/todos", handler.GetTodos)
	r.POST("/todos/deleteByID", handler.DeleteTodoByID)
	r.POST("/todos/UpdateByID", handler.UpdateTodoByID)

	r.Run(":8080") // Chạy server tại localhost:8080

}
