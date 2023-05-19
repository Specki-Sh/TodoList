package model

type Storage interface {
	Add(item Task) (int, error)
	Remove(id int) error
	GetByID(id int) (Task, error)
	GetAll() ([]Task, error)
	Update(item Task) error
}
