package service

import (
	"todo-api/models"
	"todo-api/repository"
)

func AddTodo(todo models.Todo) (models.Todo, error) {
	err := repository.CreateTodo(&todo)
	return todo, err
}

func GetTodo(todo models.Todo) ([]models.Todo, error) {
	var newTodos []models.Todo
	newTodos, err := repository.GetTodos()
	return newTodos, err
}

func DeleteTodoByID_(ID uint) error {
	err := repository.DeleteTodoByID(ID)
	return err
}

func UpdateTodoByID_(todo models.Todo) (models.Todo, error) {
	err := repository.UpdateTodoByID(&todo)
	return todo, err
}
