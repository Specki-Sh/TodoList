package controller

import (
	"todolist/domain/model"
)

func NewTaskController(storage model.TaskStorage) *TaskController {
	return &TaskController{storage: storage}
}

type TaskController struct {
	storage model.TaskStorage
}

func (t *TaskController) AddTask(task model.Task) (int, error) {
	return t.storage.Create(task)
}

func (t *TaskController) ShowCompleted() ([]model.Task, error) {
	return t.storage.SelectAllCompleted()
}

func (t *TaskController) ShowAll() ([]model.Task, error) {
	return t.storage.SelectAll()
}

func (t *TaskController) Show(id int) (model.Task, error) {
	return t.storage.SelectByID(id)
}

func (t *TaskController) Remove(id int) error {
	return t.storage.Delete(id)
}

func (t *TaskController) MarkComplete(id int) error {
	task, err := t.storage.SelectByID(id)
	if err != nil {
		return err
	}
	task.Completed = true
	return t.storage.Update(task)
}

func (t *TaskController) MarkNotComplate(id int) error {
	task, err := t.storage.SelectByID(id)
	if err != nil {
		return err
	}
	task.Completed = false
	return t.storage.Update(task)
}

func (t *TaskController) MarkAllComplete() error {
	return t.storage.MarkAllComplete()
}
