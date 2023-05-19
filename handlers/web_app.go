package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"todolist/domain/model"
	"todolist/service"

	"github.com/gin-gonic/gin"
)

type TaskData struct {
	Name string
}

func NewWebApp(todoList *service.TodoList) *WebApp {
	router := gin.Default()
	return &WebApp{
		todoList: todoList,
		router:   router,
	}
}

type WebApp struct {
	todoList *service.TodoList
	router   *gin.Engine
}

func (w *WebApp) HandleMarkAllComplete(c *gin.Context) {
	err := w.todoList.MarkAllComplete()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (w *WebApp) Route() *gin.Engine {
	t := w.router.Group("/tasks")
	t.POST("/", w.HandleAdd)
	t.GET("/", w.HandleShowAll)
	t.PUT("/", w.HandleMarkAllComplete)
	t.GET("/doned", w.HandleShowDone)
	t.PUT("/:id", validateIDParam, w.handleTasksPut)
	t.DELETE("/:id", validateIDParam, w.HandleRemove)
	return w.router
}

func (w *WebApp) HandleMarkComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := w.todoList.MarkComplete(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (w *WebApp) HandleMarkNotComplete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := w.todoList.MarkNotComplate(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (w *WebApp) HandleRemove(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "invalid product id",
		})
		return
	}

	if err := w.todoList.Remove(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func (w *WebApp) HandleShowAll(c *gin.Context) {
	allTasks, err := w.todoList.Show()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

func (w *WebApp) HandleShowDone(c *gin.Context) {
	doneTasks, err := w.todoList.ShowDoned()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, doneTasks)
}

func (w *WebApp) HandleAdd(c *gin.Context) {
	var data TaskData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "Unable to read request body")
		return
	}

	id, err := w.todoList.AddTask(model.Task{Name: data.Name})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, id)
}

func (w *WebApp) handleTasksPut(c *gin.Context) {
	complete := c.Query("complete")
	if complete == "true" {
		w.HandleMarkComplete(c)
	} else if complete == "false" {
		w.HandleMarkNotComplete(c)
	} else {
		c.String(http.StatusBadRequest, "Missing or invalid 'complete' query parameter")
	}
}

func validateIDParam(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		matched, _ := regexp.MatchString(`^\d+$`, id)
		if !matched {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"reason": "invalid id",
			})
			return
		}
	}
	c.Next()
}
