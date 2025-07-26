package dto

import "todo-app/models"

type CreateTodoRequest struct {
	Title       string          `json:"title" validate:"required,min=1,max=100"`
	Description *string         `json:"description" validate:"omitempty,max=500"`
	Priority    models.Priority `json:"priority" validate:"omitempty,oneof=LOW MEDIUM HIGH"`
}

type UpdateTodoRequest struct {
	Title       *string          `json:"title" validate:"omitempty,min=1,max=100"`
	Description *string          `json:"description" validate:"omitempty,max=500"`
	Completed   *bool            `json:"completed"`
	Priority    *models.Priority `json:"priority" validate:"omitempty,oneof=LOW MEDIUM HIGH"`
}
