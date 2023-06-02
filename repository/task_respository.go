package repository

import (
	"database/sql"
	"todolist/db"
	"todolist/domain/entity"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (s *TaskRepository) Create(item entity.Task) (int, error) {
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

func (s *TaskRepository) SelectByID(id int) (entity.Task, error) {
	var task entity.Task
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
		return entity.Task{}, err
	}
	return task, nil
}

func (s *TaskRepository) SelectAll() ([]entity.Task, error) {
	rows, err := s.db.Query(db.SelectAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
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

func (s *TaskRepository) Update(item entity.Task) error {
	_, err := s.db.Exec(db.UpdateByIDTask,
		item.UserID, item.Title, item.Description, item.DueDate, item.Priority, item.Completed, item.ID)
	return err
}

func (s *TaskRepository) SelectAllCompleted() ([]entity.Task, error) {
	rows, err := s.db.Query(db.SelectAllCompletedTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
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

func (s *TaskRepository) ReassignUser(taskID int, newUserID int) (entity.Task, error) {
	var task entity.Task

	err := s.db.QueryRow(db.SelectReassingUserTask, taskID, newUserID).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Priority, &task.Completed, &task.UserID)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (s *TaskRepository) SelectAllByUserID(userID int) ([]entity.Task, error) {
	rows, err := s.db.Query(db.SelectAllByUserIDTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
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

func (s *TaskRepository) SelectAllCompletedByUserID(userID int) ([]entity.Task, error) {
	rows, err := s.db.Query(db.SelectAllCompletedByUserIDTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
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
