package roadmapcontrollers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
	schemas "github.com/HublastX/HubLast-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

// UpdateRoadmap godoc
// @Summary Update roadmap information
// @Description Update an existing roadmap (admin only)
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param id path int true "Roadmap ID"
// @Param roadmap body CreateRoadmapRequest true "Updated roadmap information"
// @Success 200 {object} map[string]interface{} "Roadmap updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - Admin only"
// @Failure 404 {object} map[string]string "Roadmap not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/roadmaps/{id} [put]
func (c *RoadmapController) UpdateRoadmap(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roadmap ID"})
		return
	}

	var req schemas.CreateRoadmapRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update roadmaps"})
		return
	}

	roadmap, err := c.roadmapService.GetRoadmapByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Roadmap not found"})
		return
	}

	roadmap.Title = req.Title
	roadmap.Area = req.Area
	roadmap.Difficulty = req.Difficulty
	roadmap.Content = req.Content
	roadmap.CourseLinks = req.CourseLinks

	err = c.roadmapService.UpdateRoadmap(roadmap)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Roadmap updated successfully",
		"roadmap": roadmap,
	})
}
