package db

import (
	"todolist/db/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Создание таблицы users
	err := db.AutoMigrate(&model.UserModel{})
	if err != nil {
		return err
	}

	// Создание таблицы tasks
	err = db.AutoMigrate(&model.TaskModel{})
	if err != nil {
		return err
	}

	return nil
}
