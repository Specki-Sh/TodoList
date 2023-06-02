package service

import (
	"todolist/domain/entity"
	u "todolist/domain/use_cases"
)

func NewUserService(storage u.UserRepository) *UserService {
	return &UserService{storage: storage}
}

type UserService struct {
	storage u.UserRepository
}

func (u *UserService) Add(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return u.storage.Create(user)
}

func (u *UserService) Edit(user entity.User) error {
	return u.storage.Update(user)
}

func (u *UserService) Remove(id int) error {
	return u.storage.Delete(id)
}

func (u *UserService) Show(id int) (entity.User, error) {
	return u.storage.SelectByID(id)
}

func (u *UserService) ShowAll() ([]entity.User, error) {
	return u.storage.SelectAll()
}

func (u *UserService) ShowByEmailAndPassword(email string, password string) (entity.User, error) {
	return u.storage.SelectByEmailAndPassword(email, password)
}
