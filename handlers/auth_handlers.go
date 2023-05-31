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

func NewAuthHandler(authService u.AuthService) *AuthHandlers {
	return &AuthHandlers{authService: authService}
}

type AuthHandlers struct {
	authService u.AuthService
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

	userId, role, err := a.authService.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	c.Set(userCtx, userId)
	c.Set(userRoleCtx, role)
}

func (a AuthHandlers) getUserId(c *gin.Context) (int, error) {
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

func (a AuthHandlers) getUserRole(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("user role not found")
	}

	role, ok := id.(string)
	if !ok {
		return "", errors.New("user role is of invalid type")
	}

	return role, nil
}

func (a *AuthHandlers) IdentifyUserRole(c *gin.Context) {
	//id, _ := getUserId(c)
	role, _ := a.getUserRole(c)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"reason": "access denied"})
		return
	}

}
