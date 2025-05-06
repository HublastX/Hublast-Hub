package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/gin-gonic/gin"
)

// RemoveUserFromProject godoc
// @Summary Remove user from project
// @Description Removes a user from a specific project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param user body UserIDRequest true "User ID to remove"
// @Success 200 {object} map[string]string "User removed from project successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/users [delete]
func (c *ProjectController) RemoveUserFromProject(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can remove users from projects"})
		return
	}

	err = c.projectService.RemoveUserFromProject(uint(projectID), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User removed from project successfully"})
}
