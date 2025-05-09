package projectcontrollers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
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
// @Param project body schemas.CreateProjectRequest true "Updated project details"
// @Success 200 {object} map[string]interface{} "Returns message and updated project"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 404 {object} map[string]string "Project not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/projects/{id} [put]
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
	project.EstimatedTime = req.EstimatedTime
	project.DeliveryDate = req.DeliveryDate
	project.QuantyMaxUsers = req.QuantyMaxUsers

	if req.Level != "" {
		project.Level = models.ProjectLevel(req.Level)
	}

	// Precisamos implementar um método no serviço para limpar as tecnologias existentes
	if err := c.projectService.ClearProjectTechnologies(project.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear existing technologies: " + err.Error()})
		return
	}

	// Adicionar novas tecnologias frontend usando o serviço
	for _, techName := range req.FrontendTechs {
		if err := c.projectService.AddFrontendTechToProject(project.ID, techName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add frontend technology: " + err.Error()})
			return
		}
	}

	// Adicionar novas tecnologias backend usando o serviço
	for _, techName := range req.BackendTechs {
		if err := c.projectService.AddBackendTechToProject(project.ID, techName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add backend technology: " + err.Error()})
			return
		}
	}

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
