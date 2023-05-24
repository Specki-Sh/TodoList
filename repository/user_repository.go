package repository

import (
	"database/sql"
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
	err := s.db.QueryRow(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, item.Name, item.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserRepository) Update(item model.User) error {
	_, err := s.db.Exec(`UPDATE users SET name = $1, email = $2 WHERE id = $3`, item.Name, item.Email, item.ID)
	return err
}

func (s *UserRepository) Delete(id int) error {
	_, err := s.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func (s *UserRepository) SelectByID(id int) (model.User, error) {
	var user model.User
	err := s.db.QueryRow(`SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return model.User{}, err
	}
	rows, err := s.db.Query(`SELECT id, user_id, title ,description, due_date, priority, completed FROM tasks WHERE user_id = $1`, id)
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
	rows, err := s.db.Query(`SELECT u.id, u.name, u.email, t.id, t.title, t.description, t.due_date, t.priority, t.completed FROM users u LEFT JOIN tasks t ON u.id = t.user_id`)
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

		err := rows.Scan(&userID, &user.Name, &user.Email,
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
