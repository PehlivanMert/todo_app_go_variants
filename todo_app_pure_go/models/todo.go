package models

import (
	"time"
)

type Priority string

const (
	LOW    Priority = "LOW"
	MEDIUM Priority = "MEDIUM"
	HIGH   Priority = "HIGH"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null;size:100"`
	Description *string   `json:"description" gorm:"size:500"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	Priority    Priority  `json:"priority" gorm:"type:varchar(10);default:'MEDIUM'"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (t *Todo) TableName() string {
	return "todos"
}
