package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func NewWebApp(taskHandlers *TaskHandlers, userHandlers *UserHandlers, authHandlers *AuthHandlers) *WebApp {
	router := gin.Default()
	return &WebApp{
		router:       router,
		taskHandlers: *taskHandlers,
		userHandlers: *userHandlers,
		authHandlers: *authHandlers,
	}
}

type WebApp struct {
	router       *gin.Engine
	taskHandlers TaskHandlers
	userHandlers UserHandlers
	authHandlers AuthHandlers
}

func (w *WebApp) SetupRoutes() *gin.Engine {
	// Task - user
	t := w.router.Group("/tasks", w.authHandlers.userIdentity)
	{
		t.POST("/", w.taskHandlers.CreateUsed)
		t.GET("/", w.taskHandlers.GetAllUser)
		t.GET("/completed", w.taskHandlers.GetCompleted)
	}
	ta := w.router.Group("/tasks", w.authHandlers.userIdentity, w.authHandlers.AdminPermissionMiddleware())
	{
		ta.PATCH("/:id/reassign", validateIDParam, w.taskHandlers.PatchUserReassing)
	}
	tu := w.router.Group("/tasks", w.authHandlers.userIdentity, w.authHandlers.AdminPermissionMiddleware(), w.taskHandlers.TaskPermissionMiddleware())
	{
		tu.GET("/:id", validateIDParam, w.taskHandlers.GetByID)
		tu.PATCH("/:id", validateIDParam, w.taskHandlers.PatchCompeteStatus)
		tu.DELETE("/:id", validateIDParam, w.taskHandlers.Delete)
	}

	// User
	u := w.router.Group("/users", w.authHandlers.AdminPermissionMiddleware())
	{
		u.GET("/", w.userHandlers.GetAll)
		u.GET("/:id", validateIDParam, w.userHandlers.GetByID)
		u.PUT("/:id", validateIDParam, w.userHandlers.Update)
		u.DELETE("/:id", validateIDParam, w.userHandlers.Delete)
	}

	// auth
	auth := w.router.Group("/auth")
	{
		auth.POST("/sign-up", w.userHandlers.SignUp)
		auth.POST("/sign-in", w.authHandlers.SignIn)
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

func RoleMiddleware(adminHandler gin.HandlerFunc, userHandler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch role, _ := GetUserRole(c); role {
		case "admin":
			adminHandler(c)
		case "user":
			userHandler(c)
		default:
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
