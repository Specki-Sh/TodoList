package db

import (
	"fmt"
)

var dropTables = []string{
	dropReasingUserFunction,
	dropTasksTable,
	dropUsersTable,
}

func Down() error {
	for i, table := range dropTables {
		_, err := GetDBConn().Exec(table)
		if err != nil {
			return fmt.Errorf("error occurred while dropping table №%d, error is: %s", i, err.Error())
		}
	}
	return nil
}
