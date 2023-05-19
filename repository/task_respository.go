package repository

import (
	"database/sql"
	"errors"
	"todolist/domain/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (s *TaskRepository) Add(item model.Task) (int, error) {
	query := `INSERT INTO tasks (name, is_done) VALUES ($1, $2) RETURNING id`
	err := s.db.QueryRow(query, item.Name, item.IsDone).Scan(&item.ID)
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (s *TaskRepository) Remove(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (s *TaskRepository) GetByID(id int) (model.Task, error) {
	var t model.Task
	query := `SELECT id, name, is_done FROM tasks WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(&t.ID, &t.Name, &t.IsDone)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}

func (s *TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	query := `SELECT id, name, is_done FROM tasks`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t model.Task
		err = rows.Scan(&t.ID, &t.Name, &t.IsDone)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskRepository) Update(item model.Task) error {
	query := `UPDATE tasks SET name = $1, is_done = $2 WHERE id = $3`
	res, err := s.db.Exec(query, item.Name, item.IsDone, item.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
