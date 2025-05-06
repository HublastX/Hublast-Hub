package usercontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary Update user information
// @Description Update user details by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string "User updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "User not found"
// @Security BearerAuth
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Not implemented yet"})
}
