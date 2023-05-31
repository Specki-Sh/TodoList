package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func NewWebApp(taskHandlers *TaskHandlers, userHandlers *UserHandlers) *WebApp {
	router := gin.Default()
	return &WebApp{
		router:       router,
		taskHandlers: *taskHandlers,
		userHandlers: *userHandlers,
	}
}

type WebApp struct {
	router       *gin.Engine
	taskHandlers TaskHandlers
	userHandlers UserHandlers
}

func (w *WebApp) SetupRoutes() *gin.Engine {
	// Task
	t := w.router.Group("/tasks")
	{
		t.POST("/", w.taskHandlers.Create)
		t.GET("/", w.taskHandlers.GetAll)
		t.GET("/:id", validateIDParam, w.taskHandlers.GetByID)
		t.GET("/completed", w.taskHandlers.GetCompleted)
		t.PUT("/", w.taskHandlers.MarkAllComplete)
		t.PATCH("/:id", validateIDParam, w.taskHandlers.PatchCompeteStatus)
		t.PATCH("/:id/reassign", validateIDParam, w.taskHandlers.PatchUserReassing)
		t.DELETE("/:id", validateIDParam, w.taskHandlers.Delete)
	}

	// User
	u := w.router.Group("/users")
	{
		u.POST("/", w.userHandlers.Create)
		u.GET("/", w.userHandlers.GetAll)
		u.GET("/:id", validateIDParam, w.userHandlers.GetByID)
		u.PUT("/:id", validateIDParam, w.userHandlers.Update)
		u.DELETE("/:id", validateIDParam, w.userHandlers.Delete)
	}
	return w.router
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
