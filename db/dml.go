package db

// CREATE
const (
	createUsersTable = `CREATE TABLE IF NOT EXISTS users 
	(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password_hash VARCHAR NOT NULL,
		role TEXT CHECK (role IN ('user', 'admin'))
	);`
	createTasksTable = `CREATE TABLE IF NOT EXISTS tasks 
	(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		due_date TIMESTAMP,
		priority INTEGER,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
	);`
	createReasingUserFunction = `CREATE OR REPLACE FUNCTION reassign_user(task_id INTEGER, new_user_id INTEGER)
	RETURNS tasks AS $$
	DECLARE
		updated_task tasks;
	BEGIN
		UPDATE tasks SET user_id = new_user_id WHERE id = task_id RETURNING * INTO updated_task;
		RETURN updated_task;
	END;
	$$ LANGUAGE plpgsql;`
)

// DROP
const (
	dropTasksTable          = `DROP TABLE IF EXISTS tasks;`
	dropUsersTable          = `DROP TABLE IF EXISTS users;`
	dropReasingUserFunction = `DROP FUNCTION IF EXISTS reassign_user;`
)
