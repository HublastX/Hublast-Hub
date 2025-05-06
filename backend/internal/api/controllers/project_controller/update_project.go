package projectcontrollers

import (
	"net/http"
	"strconv"

	schemas "github.com/HublastX/HubLast-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

// UpdateProject godoc
// @Summary Update a project
// @Description Updates an existing project with the provided details
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body CreateProjectRequest true "Updated project details"
// @Success 200 {object} map[string]interface{} "Returns message and updated project"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 404 {object} map[string]string "Project not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id} [put]
func (c *ProjectController) UpdateProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req schemas.CreateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := c.projectService.GetProjectByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	project.Title = req.Title
	project.Description = req.Description
	project.FrontendTech = req.FrontendTech
	project.BackendTech = req.BackendTech
	project.EstimatedTime = req.EstimatedTime
	project.DeliveryDate = req.DeliveryDate

	err = c.projectService.UpdateProject(project)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project updated successfully",
		"project": project,
	})
}
