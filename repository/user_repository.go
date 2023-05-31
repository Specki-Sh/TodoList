package repository

import (
	"database/sql"
	"todolist/db"
	"todolist/domain/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) Create(item model.User) (int, error) {
	var id int
	err := s.db.QueryRow(db.InsertUser, item.Name, item.Email, item.PasswordHash, item.Role).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserRepository) Update(item model.User) error {
	_, err := s.db.Exec(db.UpdateByIDUser, item.Name, item.Email, item.PasswordHash, item.Role, item.ID)
	return err
}

func (s *UserRepository) Delete(id int) error {
	_, err := s.db.Exec(db.DeleteByIDUser, id)
	return err
}

func (s *UserRepository) SelectByID(id int) (model.User, error) {
	var user model.User
	err := s.db.QueryRow(db.SelectByIDUser, id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return model.User{}, err
	}
	rows, err := s.db.Query(db.SelectByUserIDTasks, id)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

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
			return model.User{}, err
		}
		user.Tasks = append(user.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserRepository) SelectAll() ([]model.User, error) {
	rows, err := s.db.Query(db.SelectAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersMap := make(map[int]*model.User)
	for rows.Next() {
		var (
			userID int
			taskID sql.NullInt64
			user   model.User
			task   model.Task
		)

		err := rows.Scan(&userID, &user.Name, &user.Email, &user.PasswordHash, &user.Role,
			&taskID,
			&sql.NullString{String: task.Title},
			&sql.NullString{String: task.Description},
			&sql.NullTime{Time: task.DueDate},
			&sql.NullInt64{Int64: int64(task.Priority)},
			&sql.NullBool{Bool: task.Completed})
		if err != nil {
			return nil, err
		}

		if _, ok := usersMap[userID]; !ok {
			user.ID = userID
			usersMap[userID] = &user
		}

		if taskID.Valid {
			task.ID = int(taskID.Int64)
			usersMap[userID].Tasks = append(usersMap[userID].Tasks, task)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var users []model.User
	for _, user := range usersMap {
		users = append(users, *user)
	}
	return users, nil
}

func (s *UserRepository) SelectByEmailAndPassword(email string, password string) (model.User, error) {
	var user model.User
	err := s.db.QueryRow(db.SelectByEmailAndPasswordUser, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return model.User{}, err
	}
	rows, err := s.db.Query(db.SelectByUserIDTasks, user.ID)
	if err != nil {
		return model.User{}, err
	}
	defer rows.Close()

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
			return model.User{}, err
		}
		user.Tasks = append(user.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return model.User{}, err
	}
	return user, nil
}
