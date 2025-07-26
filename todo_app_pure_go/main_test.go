package main

import (
	"testing"
	"todo-app/models"
)

func TestPriorityValidation(t *testing.T) {
	tests := []struct {
		name     string
		priority models.Priority
		valid    bool
	}{
		{"Low Priority", models.LOW, true},
		{"Medium Priority", models.MEDIUM, true},
		{"High Priority", models.HIGH, true},
		{"Invalid Priority", "INVALID", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.priority == models.LOW || tt.priority == models.MEDIUM || tt.priority == models.HIGH
			if isValid != tt.valid {
				t.Errorf("Priority validation failed for %s: expected %v, got %v", tt.priority, tt.valid, isValid)
			}
		})
	}
}

func TestTodoModel(t *testing.T) {
	description := "Test description"
	todo := models.Todo{
		Title:       "Test Todo",
		Description: &description,
		Priority:    models.HIGH,
		Completed:   false,
	}

	if todo.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got '%s'", todo.Title)
	}

	if *todo.Description != "Test description" {
		t.Errorf("Expected description 'Test description', got '%s'", *todo.Description)
	}

	if todo.Priority != models.HIGH {
		t.Errorf("Expected priority HIGH, got %s", todo.Priority)
	}

	if todo.Completed {
		t.Error("Expected completed to be false")
	}
}
