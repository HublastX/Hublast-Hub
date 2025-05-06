package roadmapcontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRoadmap godoc
// @Summary Get roadmap by ID
// @Description Get detailed information about a specific roadmap
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param id path int true "Roadmap ID"
// @Success 200 {object} map[string]interface{} "Roadmap details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Roadmap not found"
// @Security BearerAuth
// @Router /roadmaps/{id} [get]
func (c *RoadmapController) GetRoadmap(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roadmap ID"})
		return
	}

	roadmap, err := c.roadmapService.GetRoadmapByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Roadmap not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roadmap": roadmap})
}
