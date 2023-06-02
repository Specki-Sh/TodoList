package repository

import (
	"todolist/db/model"
	"todolist/domain/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) Create(item entity.User) (int, error) {
	userModel := model.NewUserModelFromEntity(item)
	result := s.db.Create(&userModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(userModel.ID), nil
}

func (s *UserRepository) Update(item entity.User) error {
	userModel := model.NewUserModelFromEntity(item)
	result := s.db.Save(&userModel)
	return result.Error
}

func (s *UserRepository) Delete(id int) error {
	result := s.db.Delete(&model.UserModel{}, id)
	return result.Error
}

func (s *UserRepository) SelectByID(id int) (entity.User, error) {
	var userModel model.UserModel
	result := s.db.Preload("Tasks").First(&userModel, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	userEntity := userModel.ToEntity()
	for _, taskModel := range userModel.Tasks {
		userEntity.Tasks = append(userEntity.Tasks, taskModel.ToEntity())
	}
	return userEntity, nil
}

func (s *UserRepository) SelectAll() ([]entity.User, error) {
	var userModels []model.UserModel
	result := s.db.Preload("Tasks").Find(&userModels)
	if result.Error != nil {
		return nil, result.Error
	}
	users := make([]entity.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = userModel.ToEntity()
		for _, taskModel := range userModel.Tasks {
			users[i].Tasks = append(users[i].Tasks, taskModel.ToEntity())
		}
	}
	return users, nil
}

func (s *UserRepository) SelectByEmailAndPassword(email string, password string) (entity.User, error) {
	var userModel model.UserModel
	result := s.db.Where("email = ? AND password_hash = ?", email, password).Preload("Tasks").First(&userModel)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	userEntity := userModel.ToEntity()
	for _, taskModel := range userModel.Tasks {
		userEntity.Tasks = append(userEntity.Tasks, taskModel.ToEntity())
	}
	return userEntity, nil
}
