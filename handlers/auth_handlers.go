package handlers

import (
	"errors"
	"net/http"
	"strings"
	u "todolist/domain/use_cases"
	"todolist/logs"

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
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := a.authUseCase.Authenticate(input.Email, input.Password)
	if err != nil {
		logs.GetLogger().Errorf("Failed to authenticate: %v", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	logs.GetLogger().Infof("User %s signed in successfully", input.Email)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (a *AuthHandlers) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logs.GetLogger().Error("Empty auth header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		logs.GetLogger().Error("Invalid auth header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid auth header"})
		return
	}

	if len(headerParts[1]) == 0 {
		logs.GetLogger().Error("Token is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "token is empty"})
		return
	}

	logs.GetLogger().Infof("Validating token: %s", headerParts[1])

	userId, role, err := a.authUseCase.ParseToken(headerParts[1])
	if err != nil {
		logs.GetLogger().Errorf("Failed to parse token: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}

	logs.GetLogger().Infof("Token validated for user ID %d with role %s", userId, role)

	c.Set(userCtx, userId)
	c.Set(userRoleCtx, role)
}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		err := errors.New("user id not found")
		logs.GetLogger().Error(err.Error())
		return 0, err
	}

	idInt, ok := id.(int)
	if !ok {
		err := errors.New("user id is of invalid type")
		logs.GetLogger().Error(err.Error())
		return 0, err
	}

	return idInt, nil
}

func GetUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(userRoleCtx)
	if !ok {
		err := errors.New("user role not found")
		logs.GetLogger().Error(err.Error())
		return "", err
	}

	roleString, ok := role.(string)
	if !ok {
		err := errors.New("user role is of invalid type")
		logs.GetLogger().Error(err.Error())
		return "", err
	}

	return roleString, nil
}

func (a *AuthHandlers) AdminPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := GetUserRole(c)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if role != "admin" {
			logs.GetLogger().Info("access forbidden")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access forbidden"})
			return
		}

		c.Next()
	}
}
