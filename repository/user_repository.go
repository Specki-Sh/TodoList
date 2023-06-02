package repository

import (
	"database/sql"
	"todolist/db"
	"todolist/domain/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) Create(item entity.User) (int, error) {
	var id int
	err := s.db.QueryRow(db.InsertUser, item.Name, item.Email, item.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserRepository) Update(item entity.User) error {
	_, err := s.db.Exec(db.UpdateByIDUser, item.Name, item.Email, item.Password, item.Role, item.ID)
	return err
}

func (s *UserRepository) Delete(id int) error {
	_, err := s.db.Exec(db.DeleteByIDUser, id)
	return err
}

func (s *UserRepository) SelectByID(id int) (entity.User, error) {
	var user entity.User
	err := s.db.QueryRow(db.SelectByIDUser, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}
	rows, err := s.db.Query(db.SelectByUserIDTasks, id)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

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
			return entity.User{}, err
		}
		user.Tasks = append(user.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *UserRepository) SelectAll() ([]entity.User, error) {
	rows, err := s.db.Query(db.SelectAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersMap := make(map[int]*entity.User)
	for rows.Next() {
		var (
			userID int
			taskID sql.NullInt64
			user   entity.User
			task   entity.Task
		)

		err := rows.Scan(&userID, &user.Name, &user.Email, &user.Password, &user.Role,
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

	var users []entity.User
	for _, user := range usersMap {
		users = append(users, *user)
	}
	return users, nil
}

func (s *UserRepository) SelectByEmailAndPassword(email string, password string) (entity.User, error) {
	var user entity.User
	err := s.db.QueryRow(db.SelectByEmailAndPasswordUser, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return entity.User{}, err
	}
	rows, err := s.db.Query(db.SelectByUserIDTasks, user.ID)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

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
			return entity.User{}, err
		}
		user.Tasks = append(user.Tasks, task)
	}
	if err := rows.Err(); err != nil {
		return entity.User{}, err
	}
	return user, nil
}
