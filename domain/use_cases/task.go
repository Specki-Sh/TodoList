package use_cases

import "todolist/domain/entity"

type TaskUseCase interface {
	AddTask(task entity.Task) (int, error)
	ShowCompleted() ([]entity.Task, error)
	ShowCompletedByUserID(userID int) ([]entity.Task, error)
	ShowAll() ([]entity.Task, error)
	ShowAllByUserID(userId int) ([]entity.Task, error)
	Show(id int) (entity.Task, error)
	Remove(id int) error
	MarkComplete(id int) error
	MarkNotComplate(id int) error
	MarkAllComplete() error
	ReassignUser(taskID int, newUserID int) (entity.Task, error)
	IsTaskAssignedToUser(userID int, taskID int) (bool, error)
}
