package repository

import (
	"errors"
	"todo-app/models"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (r *TodoRepositoryImpl) Create(todo *models.Todo) (*models.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) GetByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepositoryImpl) GetAll(completed *bool, priority *models.Priority, limit, offset int) ([]*models.Todo, error) {
	var todos []*models.Todo
	query := r.db.Model(&models.Todo{})

	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	if priority != nil {
		query = query.Where("priority = ?", *priority)
	}

	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepositoryImpl) Update(id uint, todo *models.Todo) (*models.Todo, error) {
	var existingTodo models.Todo
	if err := r.db.First(&existingTodo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	if err := r.db.Model(&existingTodo).Updates(todo).Error; err != nil {
		return nil, err
	}

	return &existingTodo, nil
}

func (r *TodoRepositoryImpl) Delete(id uint) error {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("todo not found")
		}
		return err
	}

	return r.db.Delete(&todo).Error
}

func (r *TodoRepositoryImpl) ToggleComplete(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	todo.Completed = !todo.Completed
	if err := r.db.Save(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepositoryImpl) GetTotalCount(completed *bool, priority *models.Priority) (int64, error) {
	var count int64
	query := r.db.Model(&models.Todo{})

	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	if priority != nil {
		query = query.Where("priority = ?", *priority)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
