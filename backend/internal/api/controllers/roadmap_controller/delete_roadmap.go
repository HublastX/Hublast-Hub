package roadmapcontrollers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/gin-gonic/gin"
)

// DeleteRoadmap godoc
// @Summary Delete roadmap
// @Description Delete an existing roadmap (admin only)
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param id path int true "Roadmap ID"
// @Success 200 {object} map[string]string "Roadmap deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - Admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/roadmaps/{id} [delete]
func (c *RoadmapController) DeleteRoadmap(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roadmap ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete roadmaps"})
		return
	}

	err = c.roadmapService.DeleteRoadmap(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Roadmap deleted successfully"})
}
