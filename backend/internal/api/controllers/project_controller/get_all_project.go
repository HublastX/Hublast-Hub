package projectcontrollers

import (
	"net/http"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAllProjects godoc
// @Summary Get all projects
// @Description Retrieves all projects, optionally filtered by status
// @Tags projects
// @Accept json
// @Produce json
// @Param status query string false "Filter projects by status"
// @Success 200 {object} map[string]interface{} "Returns list of projects"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /projects [get]
func (c *ProjectController) GetAllProjects(ctx *gin.Context) {

	var projects []models.Project
	var err error

	projects, err = c.projectService.GetAllProjects()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}
