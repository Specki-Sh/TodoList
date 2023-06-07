package model

import (
	"time"
	"todolist/domain/entity"
)

type TaskModel struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	DueDate     *time.Time
	Priority    int
	Completed   bool      `gorm:"default:false"`
	UserID      uint      `gorm:"not null"`
	User        UserModel `gorm:"constraint:OnDelete:CASCADE;"`
}

func (m *TaskModel) ToEntity() entity.Task {
	return entity.Task{
		ID:          int(m.ID),
		UserID:      int(m.UserID),
		Title:       m.Title,
		Description: m.Description,
		DueDate:     m.DueDate,
		Priority:    m.Priority,
		Completed:   m.Completed,
	}
}

func NewTaskModelFromEntity(e entity.Task) TaskModel {
	return TaskModel{
		ID:          uint(e.ID),
		Title:       e.Title,
		Description: e.Description,
		DueDate:     e.DueDate,
		Priority:    e.Priority,
		Completed:   e.Completed,
		UserID:      uint(e.UserID),
	}
}
