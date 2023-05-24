package repository

import (
	"database/sql"
	"todolist/domain/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (s *TaskRepository) Create(item model.Task) (int, error) {
	var id int
	err := s.db.QueryRow(`INSERT INTO tasks (user_id, title, description, due_date, priority, completed) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		item.UserID, item.Title, item.Description, item.DueDate, item.Priority, item.Completed).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TaskRepository) Delete(id int) error {
	_, err := s.db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	return err
}

func (s *TaskRepository) SelectByID(id int) (model.Task, error) {
	var task model.Task
	err := s.db.QueryRow(`SELECT id, user_id, title, description, due_date, priority, completed FROM tasks WHERE id = $1`, id).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.DueDate,
		&task.Priority,
		&task.Completed,
	)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (s *TaskRepository) SelectAll() ([]model.Task, error) {
	rows, err := s.db.Query(`SELECT id, user_id, title, description, due_date, priority, completed FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.Priority,
			&task.Completed,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskRepository) Update(item model.Task) error {
	_, err := s.db.Exec(`UPDATE tasks SET user_id=$1,title=$2 ,description=$3,due_date=$4,priority=$5 ,completed=$6 WHERE id=$7`,
		item.UserID, item.Title, item.Description, item.DueDate, item.Priority, item.Completed, item.ID)
	return err
}

func (s *TaskRepository) SelectAllCompleted() ([]model.Task, error) {
	rows, err := s.db.Query(`SELECT id,user_id,title ,description,due_date,priority ,completed FROM tasks WHERE completed=true`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.Priority,
			&task.Completed,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil

}

func (s *TaskRepository) MarkAllComplete() error {
	_, err := s.db.Exec(`UPDATE tasks SET completed=true`)
	return err

}
