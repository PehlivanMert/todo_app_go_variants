package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

// Todo represents a todo item
type Todo struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description *string   `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	Priority    Priority  `json:"priority" db:"priority"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Priority represents the priority level of a todo
type Priority string

const (
	LOW    Priority = "LOW"
	MEDIUM Priority = "MEDIUM"
	HIGH   Priority = "HIGH"
)

// TableName overrides the table name used by Pop.
func (t Todo) TableName() string {
	return "todos"
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (t *Todo) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var description string
	if t.Description != nil {
		description = *t.Description
	}

	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringLengthInRange{Field: t.Title, Name: "Title", Min: 1, Max: 100},
		&validators.StringLengthInRange{Field: description, Name: "Description", Max: 500},
		&validators.StringInclusion{Field: string(t.Priority), Name: "Priority", List: []string{"LOW", "MEDIUM", "HIGH"}},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
func (t *Todo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
func (t *Todo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
