package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProjectUsers godoc
// @Summary Get users of a specific project
// @Description Retrieves users of a project by project ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{} "Returns project users"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Project not found"
// @Security BearerAuth
// @Router /api/projects/{id}/users [get]
func (c *ProjectController) GetProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := c.projectService.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"project": project})
}
