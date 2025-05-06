package usercontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PromoteToAdmin godoc
// @Summary Promote user to admin
// @Description Change a regular user's role to admin
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string "User promoted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/users/{id}/promote [post]
func (c *UserController) PromoteToAdmin(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.userService.PromoteToAdmin(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}
