package handlers

import (
	"net/http"
	"strconv"
	"todolist/domain/entity"
	u "todolist/domain/use_cases"
	"todolist/logs"

	"github.com/gin-gonic/gin"
)

func NewTaskHandler(taskUseCase u.TaskUseCase) *TaskHandlers {
	return &TaskHandlers{taskUseCase: taskUseCase}
}

type TaskHandlers struct {
	taskUseCase u.TaskUseCase
}

func (t *TaskHandlers) MarkAllComplete(c *gin.Context) {
	err := t.taskUseCase.MarkAllComplete()
	if err != nil {
		logs.GetLogger().Errorf("Failed to mark all tasks complete: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Info("All tasks marked complete successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) MarkComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.GetLogger().Errorf("Invalid product id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.MarkComplete(id); err != nil {
		logs.GetLogger().Errorf("Failed to mark task complete: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("Task with ID %d marked complete successfully", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) MarkNotComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.GetLogger().Errorf("Invalid product id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.MarkNotComplate(id); err != nil {
		logs.GetLogger().Errorf("Failed to mark task not complete: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("Task with ID %d marked not complete successfully", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logs.GetLogger().Errorf("Invalid product id: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.Remove(id); err != nil {
		logs.GetLogger().Errorf("Failed to remove task: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("Task with ID %d removed successfully", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) GetAllAdmin(c *gin.Context) {
	allTasks, err := t.taskUseCase.ShowAll()
	if err != nil {
		logs.GetLogger().Errorf("Failed to show all tasks: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Info("All tasks shown successfully")

	c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandlers) GetAllUser(c *gin.Context) {
	userID, _ := GetUserId(c)
	allTasks, err := t.taskUseCase.ShowAllByUserID(userID)
	if err != nil {
		logs.GetLogger().Errorf("Failed to show all tasks by user ID: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("All tasks shown successfully for user ID %d", userID)

	c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandlers) GetCompleted(c *gin.Context) {
	userId, _ := GetUserId(c)
	doneTasks, err := t.taskUseCase.ShowCompletedByUserID(userId)
	if err != nil {
		logs.GetLogger().Errorf("Failed to show completed tasks by user ID: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("Completed tasks shown successfully for user ID %d", userId)

	c.JSON(http.StatusOK, doneTasks)
}

func (t *TaskHandlers) CreateUsed(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID, _ = GetUserId(c)
	id, err := t.taskUseCase.AddTask(task)
	if err != nil {
		logs.GetLogger().Errorf("Failed to add task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("Task with ID %d added successfully for user ID %d", id, task.UserID)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (t *TaskHandlers) CreateAdmin(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := t.taskUseCase.AddTask(task)
	if err != nil {
		logs.GetLogger().Errorf("Failed to add task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("Task with ID %d added successfully by admin", id)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (t *TaskHandlers) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logs.GetLogger().Errorf("Invalid task ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.taskUseCase.Show(id)
	if err != nil {
		logs.GetLogger().Errorf("Failed to show task with ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("Task with ID %d shown successfully", id)

	c.JSON(http.StatusOK, task)
}

func (t *TaskHandlers) PatchCompeteStatus(c *gin.Context) {
	complete := c.Query("complete")
	if complete == "true" {
		t.MarkComplete(c)
	} else if complete == "false" {
		t.MarkNotComplete(c)
	} else {
		logs.GetLogger().Error("Missing or invalid 'complete' query parameter")
		c.String(http.StatusBadRequest, "Missing or invalid 'complete' query parameter")
	}
}

func (t *TaskHandlers) PatchUserReassing(c *gin.Context) {
	taskID := c.Param("id")
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		logs.GetLogger().Errorf("Invalid task ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var body struct {
		UserID int `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.taskUseCase.ReassignUser(taskIDInt, body.UserID)
	if err != nil {
		logs.GetLogger().Errorf("Failed to reassign user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("Task with ID %d reassigned successfully to user ID %d", taskIDInt, body.UserID)

	c.JSON(http.StatusOK, task)
}

func (t *TaskHandlers) TaskPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("id")
		taskIDInt, _ := strconv.Atoi(taskID)
		userId, _ := GetUserId(c)
		isTaskAssignedToUser, err := t.taskUseCase.IsTaskAssignedToUser(userId, taskIDInt)
		if err != nil {
			logs.GetLogger().Errorf("Failed to check if task is assigned to user: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if isTaskAssignedToUser {
			c.Next()
		} else {
			logs.GetLogger().Infof("Access denied for user ID %d to task with ID %d", userId, taskIDInt)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		}
	}
}
