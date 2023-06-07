package handlers

import (
	"errors"
	"net/http"
	"strings"
	u "todolist/domain/use_cases"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userRoleCtx         = "userRole"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(authUseCase u.AuthUseCase) *AuthHandlers {
	return &AuthHandlers{authUseCase: authUseCase}
}

type AuthHandlers struct {
	authUseCase u.AuthUseCase
}

func (a *AuthHandlers) SignIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := a.authUseCase.Authenticate(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (a *AuthHandlers) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": "token is empty"})
		return
	}

	userId, role, err := a.authUseCase.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userCtx, userId)
	c.Set(userRoleCtx, role)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func GetUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(userRoleCtx)
	if !ok {
		return "", errors.New("user role not found")
	}

	roleString, ok := role.(string)
	if !ok {
		return "", errors.New("user role is of invalid type")
	}

	return roleString, nil
}

func (a *AuthHandlers) AdminPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := GetUserRole(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error", "error:": err.Error()})
			return
		}
		if role == "admin" {
			c.Set("AdminPermission", true)
			c.Next()
		} else {
			c.Next()
		}
	}
}
