package handlers

import (
	"net/http"
	"strconv"
	"todolist/domain/entity"
	u "todolist/domain/use_cases"
	"todolist/logs"

	"github.com/gin-gonic/gin"
)

func NewUserHandlers(userUseCase u.UserUseCase) *UserHandlers {
	return &UserHandlers{userUseCase: userUseCase}
}

type UserHandlers struct {
	userUseCase u.UserUseCase
}

func (u *UserHandlers) SignUp(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := u.userUseCase.Add(user)
	if err != nil {
		logs.GetLogger().Errorf("Failed to add user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("User with ID %d added successfully", id)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (u *UserHandlers) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logs.GetLogger().Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logs.GetLogger().Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := u.userUseCase.Edit(user); err != nil {
		logs.GetLogger().Errorf("Failed to edit user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("User with ID %d updated successfully", id)

	c.Status(http.StatusOK)
}

func (u *UserHandlers) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logs.GetLogger().Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := u.userUseCase.Remove(id); err != nil {
		logs.GetLogger().Errorf("Failed to remove user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("User with ID %d removed successfully", id)

	c.Status(http.StatusOK)
}

func (u *UserHandlers) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logs.GetLogger().Errorf("Invalid user ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userUseCase.Show(id)
	if err != nil {
		logs.GetLogger().Errorf("Failed to show user with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Infof("User with ID %d shown successfully", id)

	c.JSON(http.StatusOK, user)
}

func (u *UserHandlers) GetAll(c *gin.Context) {
	users, err := u.userUseCase.ShowAll()
	if err != nil {
		logs.GetLogger().Errorf("Failed to show all users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logs.GetLogger().Info("All users shown successfully")

	c.JSON(http.StatusOK, users)
}
