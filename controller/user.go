package controller

import "todolist/domain/model"

func NewUserController(storage model.UserStorage) *UserController {
	return &UserController{storage: storage}
}

type UserController struct {
	storage model.UserStorage
}

func (u *UserController) Add(user model.User) (int, error) {
	return u.storage.Create(user)
}

func (u *UserController) Edit(user model.User) error {
	return u.storage.Update(user)
}

func (u *UserController) Remove(id int) error {
	return u.storage.Delete(id)
}

func (u *UserController) Show(id int) (model.User, error) {
	return u.storage.SelectByID(id)
}

func (u *UserController) ShowAll() ([]model.User, error) {
	return u.storage.SelectAll()
}
