package roadmapcontrollers

import (
	"net/http"

	"github.com/HublastX/HubLast-Hub/internal/models"
	schemas "github.com/HublastX/HubLast-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

// CreateRoadmap godoc
// @Summary Create a new roadmap
// @Description Create a new learning roadmap (admin only)
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param roadmap body CreateRoadmapRequest true "Roadmap information"
// @Success 201 {object} map[string]interface{} "Roadmap created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - Admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/roadmaps [post]
func (c *RoadmapController) CreateRoadmap(ctx *gin.Context) {
	var req schemas.CreateRoadmapRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create roadmaps"})
		return
	}

	roadmap := &models.Roadmap{
		Title:       req.Title,
		Area:        req.Area,
		Difficulty:  req.Difficulty,
		Content:     req.Content,
		CourseLinks: req.CourseLinks,
	}

	err := c.roadmapService.CreateRoadmap(roadmap)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Roadmap created successfully",
		"roadmap": roadmap,
	})
}
