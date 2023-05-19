package service

import (
	"todolist/domain/model"
)

func NewTodoList(storage model.Storage) *TodoList {
	return &TodoList{storage: storage}
}

type TodoList struct {
	storage model.Storage
}

func (t *TodoList) AddTask(task model.Task) (int, error) {
	return t.storage.Add(task)
}

func (t *TodoList) ShowDoned() ([]model.Task, error) {
	tasks, err := t.storage.GetAll()
	if err != nil {
		return nil, err
	}
	donedTasks := make([]model.Task, 0, len(tasks))
	for _, task := range tasks {
		if task.IsDone {
			donedTasks = append(donedTasks, task)
		}
	}
	return donedTasks, nil
}

func (t *TodoList) Show() ([]model.Task, error) {
	tasks, err := t.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TodoList) Remove(id int) error {
	return t.storage.Remove(id)
}

func (t *TodoList) MarkComplete(id int) error {
	task, err := t.storage.GetByID(id)
	if err != nil {
		return err
	}
	task.IsDone = true
	return t.storage.Update(task)
}

func (t *TodoList) MarkNotComplate(id int) error {
	task, err := t.storage.GetByID(id)
	if err != nil {
		return err
	}
	task.IsDone = false
	return t.storage.Update(task)
}

func (t *TodoList) MarkAllComplete() error {
	tasks, err := t.storage.GetAll()
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if err := t.MarkComplete(task.ID); err != nil {
			return err
		}
	}
	return nil
}
