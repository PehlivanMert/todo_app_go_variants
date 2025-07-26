package repository

import (
	"todo-app/models"
)

type TodoRepository interface {
	Create(todo *models.Todo) (*models.Todo, error)
	GetByID(id uint) (*models.Todo, error)
	GetAll(completed *bool, priority *models.Priority, limit, offset int) ([]*models.Todo, error)
	Update(id uint, todo *models.Todo) (*models.Todo, error)
	Delete(id uint) error
	ToggleComplete(id uint) (*models.Todo, error)
	GetTotalCount(completed *bool, priority *models.Priority) (int64, error)
}
