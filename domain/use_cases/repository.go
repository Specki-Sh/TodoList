package use_cases

import "todolist/domain/model"

type TaskRepository interface {
	Create(item model.Task) (int, error)
	Delete(id int) error
	SelectByID(id int) (model.Task, error)
	SelectAll() ([]model.Task, error)
	SelectAllByUserID(userID int) ([]model.Task, error)
	SelectAllCompletedByUserID(userID int) ([]model.Task, error)
	Update(item model.Task) error
	SelectAllCompleted() ([]model.Task, error)
	MarkAllComplete() error
	ReassignUser(taskID int, newUserID int) (model.Task, error)
}

type UserRepository interface {
	Create(item model.User) (int, error)
	Update(item model.User) error
	Delete(id int) error
	SelectByID(id int) (model.User, error)
	SelectAll() ([]model.User, error)
	SelectByEmailAndPassword(email string, password string) (model.User, error)
}
