package use_cases

import "todolist/domain/model"

type UserUseCase interface {
	Add(user model.User) (int, error)
	Edit(user model.User) error
	Remove(id int) error
	Show(id int) (model.User, error)
	ShowAll() ([]model.User, error)
	ShowByEmailAndPassword(email string, password string) (model.User, error)
}
