package repository

import (
	"database/sql"
	"todolist/db"
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
	err := s.db.QueryRow(db.InsertTask,
		item.UserID, item.Title, item.Description, item.DueDate, item.Priority, item.Completed).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TaskRepository) Delete(id int) error {
	_, err := s.db.Exec(db.DeleteByIDTask, id)
	return err
}

func (s *TaskRepository) SelectByID(id int) (model.Task, error) {
	var task model.Task
	err := s.db.QueryRow(db.SelectByIDTask, id).Scan(
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
	rows, err := s.db.Query(db.SelectAllTasks)
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
	_, err := s.db.Exec(db.UpdateByIDTask,
		item.UserID, item.Title, item.Description, item.DueDate, item.Priority, item.Completed, item.ID)
	return err
}

func (s *TaskRepository) SelectAllCompleted() ([]model.Task, error) {
	rows, err := s.db.Query(db.SelectAllCompletedTasks)
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
	_, err := s.db.Exec(db.UpdateMakeAllCompletedTasks)
	return err
}

func (s *TaskRepository) ReassignUser(taskID int, newUserID int) (model.Task, error) {
	var task model.Task

	err := s.db.QueryRow(db.SelectReassingUserTask, taskID, newUserID).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Completed, &task.UserID)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}
