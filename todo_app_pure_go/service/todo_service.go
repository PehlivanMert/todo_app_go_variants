package service

import (
	"todo-app/dto"
	"todo-app/models"
)

type TodoService interface {
	CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error)
	GetTodoByID(id uint) (*dto.TodoResponse, error)
	GetAllTodos(completed *bool, priority *models.Priority, limit, offset int) ([]*dto.TodoResponse, int64, error)
	UpdateTodo(id uint, req *dto.UpdateTodoRequest) (*dto.TodoResponse, error)
	DeleteTodo(id uint) error
	ToggleTodoComplete(id uint) (*dto.TodoResponse, error)
}
