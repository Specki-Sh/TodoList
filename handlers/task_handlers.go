package handlers

import (
	"net/http"
	"strconv"
	"todolist/controller"
	"todolist/domain/model"

	"github.com/gin-gonic/gin"
)

func NewTaskHandler(taskController *controller.TaskController) *TaskHandlers {
	return &TaskHandlers{taskController: taskController}
}

type TaskHandlers struct {
	taskController *controller.TaskController
}

func (t *TaskHandlers) MarkAllComplete(c *gin.Context) {
	err := t.taskController.MarkAllComplete()
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

	if err := t.taskController.MarkComplete(id); err != nil {
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

	if err := t.taskController.MarkNotComplate(id); err != nil {
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

	if err := t.taskController.Remove(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (t *TaskHandlers) GetAll(c *gin.Context) {
	allTasks, err := t.taskController.ShowAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

func (t *TaskHandlers) GetCompleted(c *gin.Context) {
	doneTasks, err := t.taskController.ShowCompleted()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, doneTasks)
}

func (t *TaskHandlers) Create(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := t.taskController.AddTask(task)
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
	task, err := t.taskController.Show(id)
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

	task, err := t.taskController.ReassignUser(taskIDInt, body.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
