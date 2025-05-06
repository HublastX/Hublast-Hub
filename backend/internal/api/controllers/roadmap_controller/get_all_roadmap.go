package roadmapcontrollers

import (
	"net/http"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllRoadmaps godoc
// @Summary Get all roadmaps
// @Description Get a list of all roadmaps, optionally filtered by area
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param area query string false "Filter by roadmap area"
// @Success 200 {object} map[string]interface{} "List of roadmaps"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /roadmaps [get]
func (c *RoadmapController) GetAllRoadmaps(ctx *gin.Context) {
	area := ctx.Query("area")

	var roadmaps []models.Roadmap
	var err error

	if area != "" {
		roadmaps, err = c.roadmapService.GetRoadmapsByArea(models.RoadmapArea(area))
	} else {
		roadmaps, err = c.roadmapService.GetAllRoadmaps()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"roadmaps": roadmaps})
}
