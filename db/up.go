package db

import (
	"fmt"
)

var createTables = []string{
	createUsersTable,
	createTasksTable,
	createReasingUserFunction,
}

func Up() error {
	for i, table := range createTables {
		_, err := GetDBConn().Exec(table)
		if err != nil {
			return fmt.Errorf("error occurred while creating table №%d, error is: %s", i, err.Error())
		}
	}
	return nil
}
