package use_cases

import "todolist/domain/entity"

type TaskRepository interface {
	Create(item entity.Task) (int, error)
	Delete(id int) error
	SelectByID(id int) (entity.Task, error)
	SelectAll() ([]entity.Task, error)
	SelectAllByUserID(userID int) ([]entity.Task, error)
	SelectAllCompletedByUserID(userID int) ([]entity.Task, error)
	Update(item entity.Task) error
	SelectAllCompleted() ([]entity.Task, error)
	MarkAllComplete() error
	ReassignUser(taskID int, newUserID int) (entity.Task, error)
}

type UserRepository interface {
	Create(item entity.User) (int, error)
	Update(item entity.User) error
	Delete(id int) error
	SelectByID(id int) (entity.User, error)
	SelectAll() ([]entity.User, error)
	SelectByEmailAndPassword(email string, password string) (entity.User, error)
}
