package use_cases

import "todolist/domain/model"

type TaskUseCase interface {
	AddTask(task model.Task) (int, error)
	ShowCompleted() ([]model.Task, error)
	ShowCompletedByUserID(userID int) ([]model.Task, error)
	ShowAll() ([]model.Task, error)
	ShowAllByUserID(userId int) ([]model.Task, error)
	Show(id int) (model.Task, error)
	Remove(id int) error
	MarkComplete(id int) error
	MarkNotComplate(id int) error
	MarkAllComplete() error
	ReassignUser(taskID int, newUserID int) (model.Task, error)
	IsTaskAssignedToUser(userID int, taskID int) (bool, error)
}
