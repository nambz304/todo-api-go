package repository

import (
	"fmt"
	"todo-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "postgresql://todo_db_fz9m_user:dwTBvvliy8V2kbXwvqS6PZ6nQhZqtkH6@dpg-d59psqm3jp1c73c798gg-a.singapore-postgres.render.com/todo_db_fz9m"

	// Dùng Postgres driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	DB = db
	DB.AutoMigrate(&models.Todo{}) // Tự tạo bảng todos
	fmt.Println("Connected to database successfully!")
}

func CreateTodo(todo *models.Todo) error {
	return DB.Create(todo).Error
}

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := DB.Find(&todos).Error
	return todos, err
}

func DeleteTodoByID(ID uint) error {
	return DB.Delete(&models.Todo{}, ID).Error // Xóa theo ID
}

func UpdateTodoByID(todo *models.Todo) error {
	return DB.Save(todo).Error
}

// // Thêm hàm GetByID, Update, Delete tương tự...
// func GetByID() ([]models.Todo, error) {
// 	return
// }
