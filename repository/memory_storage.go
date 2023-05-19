package repository

import (
	"errors"
	"todolist/domain/model"
)

type MemoryStorage struct {
	tasks  []model.Task
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	task1 := model.Task{Name: "Do hm"}
	task2 := model.Task{Name: "Clean table"}
	task3 := model.Task{Name: "Sleep"}
	ms := MemoryStorage{
		tasks:  []model.Task{},
		nextID: 1,
	}
	ms.Add(task1)
	ms.Add(task2)
	ms.Add(task3)
	return &ms
}

func (s *MemoryStorage) Add(item model.Task) error {
	item.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, item)
	return nil
}

func (s *MemoryStorage) Remove(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *MemoryStorage) GetByID(id int) (model.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func (s *MemoryStorage) GetAll() ([]model.Task, error) {
	return s.tasks, nil
}

func (s *MemoryStorage) Update(item model.Task) error {
	for i, task := range s.tasks {
		if task.ID == item.ID {
			s.tasks[i] = item
			return nil
		}
	}
	return errors.New("task not found")
}
