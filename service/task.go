package service

import (
	"todolist/domain/model"
	u "todolist/domain/use_cases"
)

func NewTaskService(storage u.TaskRepository) *TaskService {
	return &TaskService{storage: storage}
}

type TaskService struct {
	storage u.TaskRepository
}

func (t *TaskService) AddTask(task model.Task) (int, error) {
	return t.storage.Create(task)
}

func (t *TaskService) ShowCompleted() ([]model.Task, error) {
	return t.storage.SelectAllCompleted()
}

func (t *TaskService) ShowAll() ([]model.Task, error) {
	return t.storage.SelectAll()
}

func (t *TaskService) Show(id int) (model.Task, error) {
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

func (t *TaskService) ReassignUser(taskID int, newUserID int) (model.Task, error) {
	return t.storage.ReassignUser(taskID, newUserID)
}
