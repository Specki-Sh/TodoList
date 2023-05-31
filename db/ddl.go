package db

// INSERT
const (
	InsertTask = `INSERT INTO tasks (user_id, title, description, due_date, priority, completed) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	InsertUser = `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
)

// SELECT
const (
	SelectByIDTask          = `SELECT id, user_id, title, description, due_date, priority, completed FROM tasks WHERE id=$1`
	SelectAllTasks          = `SELECT id, user_id, title, description, due_date, priority, completed FROM tasks`
	SelectAllCompletedTasks = `SELECT id, user_id, title, description, due_date, priority, completed FROM tasks WHERE completed=true`
	SelectReassingUserTask  = `SELECT * FROM reassign_user($1, $2)`
	SelectByUserIDTasks     = `SELECT id, user_id, title, description, due_date, priority, completed FROM tasks WHERE user_id = $1`

	SelectByIDUser = `SELECT id, name, email FROM users WHERE id = $1`
	SelectAllUsers = `SELECT u.id, u.name, u.email, t.id, t.title, t.description, t.due_date, t.priority, t.completed FROM users u LEFT JOIN tasks t ON u.id = t.user_id`
)

// UPDATE
const (
	UpdateByIDTask              = `UPDATE tasks SET user_id=$1, title=$2, description=$3, due_date=$4, priority=$5, completed=$6 WHERE id=$7`
	UpdateMakeAllCompletedTasks = `UPDATE tasks SET completed=true`

	UpdateByIDUser = `UPDATE users SET name=$1, email=$2 WHERE id=$3`
)

// DELETE
const (
	DeleteByIDTask = `DELETE FROM tasks WHERE id=$1`

	DeleteByIDUser = `DELETE FROM users WHERE id=$1`
)