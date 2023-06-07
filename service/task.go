package service

import (
	"todolist/domain/entity"
	u "todolist/domain/use_cases"
)

func NewTaskService(storage u.TaskRepository) *TaskService {
	return &TaskService{storage: storage}
}

type TaskService struct {
	storage u.TaskRepository
}

func (t *TaskService) AddTask(task entity.Task) (int, error) {
	return t.storage.Create(task)
}

func (t *TaskService) ShowCompleted() ([]entity.Task, error) {
	return t.storage.SelectAllCompleted()
}

func (t *TaskService) ShowCompletedByUserID(userID int) ([]entity.Task, error) {
	return t.storage.SelectAllCompletedByUserID(userID)
}

func (t *TaskService) ShowAll() ([]entity.Task, error) {
	return t.storage.SelectAll()
}

func (t *TaskService) ShowAllByUserID(userId int) ([]entity.Task, error) {
	return t.storage.SelectAllByUserID(userId)
}

func (t *TaskService) Show(id int) (entity.Task, error) {
	return t.storage.SelectByID(id)
}

func (t *TaskService) Remove(id int) error {
	return t.storage.Delete(id)
}

func (t *TaskService) MarkComplete(id int) error {
	task, err := t.storage.SelectByID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	return t.storage.Update(task)
}

func (t *TaskService) MarkNotComplate(id int) error {
	task, err := t.storage.SelectByID(id)
	if err != nil {
		return err
	}
	task.Completed = false
	return t.storage.Update(task)
}

func (t *TaskService) MarkAllComplete() error {
	return t.storage.MarkAllComplete()
}

func (t *TaskService) ReassignUser(taskID int, newUserID int) (entity.Task, error) {
	return t.storage.ReassignUser(taskID, newUserID)
}

func (t *TaskService) IsTaskAssignedToUser(userID int, taskID int) (bool, error) {
	task, err := t.Show(taskID)
	if err != nil {
		return false, err
	}
	return userID == task.UserID, nil
}
