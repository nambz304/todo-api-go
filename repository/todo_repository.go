package repository

import (
	"os"
	"todo-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL") // Heroku tự cung cấp env này

	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/todo_db?..." // fallback local
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db
	DB.AutoMigrate(&models.Todo{}) // Tạo bảng tự động

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
