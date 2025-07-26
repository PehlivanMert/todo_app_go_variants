package service

import (
	"errors"
	"strconv"
	"todo-app/dto"
	"todo-app/models"
	"todo-app/repository"
	"todo-app/utils"
)

type TodoServiceImpl struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (s *TodoServiceImpl) CreateTodo(req *dto.CreateTodoRequest) (*dto.TodoResponse, error) {
	// Validate request
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.New("validation failed: " + validationErrors[0])
	}

	// Set default priority if not provided
	if req.Priority == "" {
		req.Priority = models.MEDIUM
	}

	// Create todo model
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Completed:   false,
	}

	// Save to database
	createdTodo, err := s.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	return s.todoToResponse(createdTodo), nil
}

func (s *TodoServiceImpl) GetTodoByID(id uint) (*dto.TodoResponse, error) {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.todoToResponse(todo), nil
}

func (s *TodoServiceImpl) GetAllTodos(completed *bool, priority *models.Priority, limit, offset int) ([]*dto.TodoResponse, int64, error) {
	// Validate pagination parameters
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	// Get todos from repository
	todos, err := s.todoRepo.GetAll(completed, priority, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := s.todoRepo.GetTotalCount(completed, priority)
	if err != nil {
		return nil, 0, err
	}

	// Convert to response DTOs
	responses := make([]*dto.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = s.todoToResponse(todo)
	}

	return responses, total, nil
}

func (s *TodoServiceImpl) UpdateTodo(id uint, req *dto.UpdateTodoRequest) (*dto.TodoResponse, error) {
	// Validate request
	if validationErrors := utils.ValidateStruct(req); len(validationErrors) > 0 {
		return nil, errors.New("validation failed: " + validationErrors[0])
	}

	// Check if todo exists
	existingTodo, err := s.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != nil {
		existingTodo.Title = *req.Title
	}
	if req.Description != nil {
		existingTodo.Description = req.Description
	}
	if req.Completed != nil {
		existingTodo.Completed = *req.Completed
	}
	if req.Priority != nil {
		existingTodo.Priority = *req.Priority
	}

	// Save to database
	updatedTodo, err := s.todoRepo.Update(id, existingTodo)
	if err != nil {
		return nil, err
	}

	return s.todoToResponse(updatedTodo), nil
}

func (s *TodoServiceImpl) DeleteTodo(id uint) error {
	return s.todoRepo.Delete(id)
}

func (s *TodoServiceImpl) ToggleTodoComplete(id uint) (*dto.TodoResponse, error) {
	todo, err := s.todoRepo.ToggleComplete(id)
	if err != nil {
		return nil, err
	}

	return s.todoToResponse(todo), nil
}

// Helper method to convert Todo model to TodoResponse DTO
func (s *TodoServiceImpl) todoToResponse(todo *models.Todo) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		Priority:    todo.Priority,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// Helper method to parse string to uint
func parseUint(s string) (uint, error) {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(u64), nil
}
