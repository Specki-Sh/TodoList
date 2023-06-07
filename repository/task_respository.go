package repository

import (
	"todolist/db/model"
	"todolist/domain/entity"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (s *TaskRepository) Create(item entity.Task) (int, error) {
	taskModel := model.NewTaskModelFromEntity(item)
	result := s.db.Create(&taskModel)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(taskModel.ID), nil
}

func (s *TaskRepository) Delete(id int) error {
	result := s.db.Delete(&model.TaskModel{}, id)
	return result.Error
}

func (s *TaskRepository) SelectByID(id int) (entity.Task, error) {
	var taskModel model.TaskModel
	result := s.db.First(&taskModel, id)
	if result.Error != nil {
		return entity.Task{}, result.Error
	}
	return taskModel.ToEntity(), nil
}

func (s *TaskRepository) SelectAll() ([]entity.Task, error) {
	var taskModels []model.TaskModel
	result := s.db.Find(&taskModels)
	if result.Error != nil {
		return nil, result.Error
	}
	tasks := make([]entity.Task, len(taskModels))
	for i, taskModel := range taskModels {
		tasks[i] = taskModel.ToEntity()
	}
	return tasks, nil
}

func (s *TaskRepository) Update(item entity.Task) error {
	taskModel := model.NewTaskModelFromEntity(item)
	result := s.db.Save(&taskModel)
	return result.Error
}

func (s *TaskRepository) SelectAllCompleted() ([]entity.Task, error) {
	var taskModels []model.TaskModel
	result := s.db.Where("completed = ?", true).Find(&taskModels)
	if result.Error != nil {
		return nil, result.Error
	}
	tasks := make([]entity.Task, len(taskModels))
	for i, taskModel := range taskModels {
		tasks[i] = taskModel.ToEntity()
	}
	return tasks, nil
}

func (s *TaskRepository) MarkAllComplete() error {
	result := s.db.Model(&model.TaskModel{}).Where("completed = ?", false).Update("completed", true)
	return result.Error
}

func (s *TaskRepository) ReassignUser(taskID int, newUserID int) (entity.Task, error) {
	var taskModel model.TaskModel

	result := s.db.Model(&taskModel).Where("id = ?", taskID).Update("user_id", newUserID)
	if result.Error != nil {
		return entity.Task{}, result.Error
	}

	result = s.db.First(&taskModel, taskID)
	if result.Error != nil {
		return entity.Task{}, result.Error
	}

	return taskModel.ToEntity(), nil
}

func (s *TaskRepository) SelectAllByUserID(userID int) ([]entity.Task, error) {
	var taskModels []model.TaskModel
	result := s.db.Where("user_id = ?", userID).Find(&taskModels)
	if result.Error != nil {
		return nil, result.Error
	}
	tasks := make([]entity.Task, len(taskModels))
	for i, taskModel := range taskModels {
		tasks[i] = taskModel.ToEntity()
	}
	return tasks, nil
}

func (s *TaskRepository) SelectAllCompletedByUserID(userID int) ([]entity.Task, error) {
	var taskModels []model.TaskModel
	result := s.db.Where("user_id = ? AND completed = ?", userID, true).Find(&taskModels)
	if result.Error != nil {
		return nil, result.Error
	}
	tasks := make([]entity.Task, len(taskModels))
	for i, taskModel := range taskModels {
		tasks[i] = taskModel.ToEntity()
	}
	return tasks, nil
}
