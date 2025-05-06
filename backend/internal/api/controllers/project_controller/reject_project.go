package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/gin-gonic/gin"
)

// RejectProject godoc
// @Summary Reject a project
// @Description Rejects a pending project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Project rejected successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/reject [post]
func (c *ProjectController) RejectProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can reject projects"})
		return
	}

	err = c.projectService.RejectProject(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project rejected successfully"})
}
