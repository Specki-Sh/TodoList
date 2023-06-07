package use_cases

import "todolist/domain/entity"

type UserUseCase interface {
	Add(user entity.User) (int, error)
	Edit(user entity.User) error
	Remove(id int) error
	Show(id int) (entity.User, error)
	ShowAll() ([]entity.User, error)
	ShowByEmailAndPassword(email string, password string) (entity.User, error)
}
