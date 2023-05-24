package model

type TaskStorage interface {
	Create(item Task) (int, error)
	Delete(id int) error
	SelectByID(id int) (Task, error)
	SelectAll() ([]Task, error)
	Update(item Task) error
	SelectAllCompleted() ([]Task, error)
	MarkAllComplete() error
}

type UserStorage interface {
	Create(item User) (int, error)
	Update(item User) error
	Delete(id int) error
	SelectByID(id int) (User, error)
	SelectAll() ([]User, error)
}
