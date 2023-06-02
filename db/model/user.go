package model

import "todolist/domain/entity"

type UserModel struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:user;check:role IN ('user', 'admin')"`
	Tasks        []TaskModel
}

func (m *UserModel) ToEntity() entity.User {
	return entity.User{
		ID:       int(m.ID),
		Name:     m.Name,
		Email:    m.Email,
		Password: m.PasswordHash,
		Role:     m.Role,
	}
}

func NewUserModelFromEntity(e entity.User) UserModel {
	return UserModel{
		ID:           uint(e.ID),
		Name:         e.Name,
		Email:        e.Email,
		PasswordHash: e.Password,
		Role:         e.Role,
	}
}
