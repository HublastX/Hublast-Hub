package usercontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of users"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"users": response})
}
