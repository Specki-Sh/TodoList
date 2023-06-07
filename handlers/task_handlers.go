package handlers

import (
	"net/http"
	"strconv"
	"todolist/domain/entity"
	u "todolist/domain/use_cases"

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
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) MarkComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.MarkComplete(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) MarkNotComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.MarkNotComplate(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := t.taskUseCase.Remove(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) GetAllAdmin(c *gin.Context) {
	allTasks, err := t.taskUseCase.ShowAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandlers) GetAllUser(c *gin.Context) {
	userID, _ := GetUserId(c)
	allTasks, err := t.taskUseCase.ShowAllByUserID(userID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandlers) GetCompleted(c *gin.Context) {
	userId, _ := GetUserId(c)
	doneTasks, err := t.taskUseCase.ShowCompletedByUserID(userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, doneTasks)
}

func (t *TaskHandlers) CreateUsed(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.UserID, _ = GetUserId(c)
	id, err := t.taskUseCase.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (t *TaskHandlers) CreateAdmin(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := t.taskUseCase.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (t *TaskHandlers) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := t.taskUseCase.Show(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (t *TaskHandlers) PatchCompeteStatus(c *gin.Context) {
	complete := c.Query("complete")
	if complete == "true" {
		t.MarkComplete(c)
	} else if complete == "false" {
		t.MarkNotComplete(c)
	} else {
		c.String(http.StatusBadRequest, "Missing or invalid 'complete' query parameter")
	}
}

func (t *TaskHandlers) PatchUserReassing(c *gin.Context) {
	taskID := c.Param("id")
	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var body struct {
		UserID int `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := t.taskUseCase.ReassignUser(taskIDInt, body.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *TaskHandlers) TaskPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("AdminPermission") {
			c.Next()
			return
		}

		taskID := c.Param("id")
		taskIDInt, _ := strconv.Atoi(taskID)
		userId, _ := GetUserId(c)
		isTaskAssignedToUser, err := t.taskUseCase.IsTaskAssignedToUser(userId, taskIDInt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
		if isTaskAssignedToUser {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		}
	}
}
