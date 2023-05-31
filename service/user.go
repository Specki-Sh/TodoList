package service

import (
	"todolist/domain/model"
	u "todolist/domain/use_cases"
)

func NewUserService(storage u.UserRepository) *UserService {
	return &UserService{storage: storage}
}

type UserService struct {
	storage u.UserRepository
}

func (u *UserService) Add(user model.User) (int, error) {
	return u.storage.Create(user)
}

func (u *UserService) Edit(user model.User) error {
	return u.storage.Update(user)
}

func (u *UserService) Remove(id int) error {
	return u.storage.Delete(id)
}

func (u *UserService) Show(id int) (model.User, error) {
	return u.storage.SelectByID(id)
}

func (u *UserService) ShowAll() ([]model.User, error) {
	return u.storage.SelectAll()
}

func (u *UserService) ShowByEmailAndPassword(email string, password string) (model.User, error) {
	return u.storage.SelectByEmailAndPassword(email, password)
}
