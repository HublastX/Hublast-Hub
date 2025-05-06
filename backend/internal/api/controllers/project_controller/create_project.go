package projectcontrollers

import (
	"net/http"

	"github.com/HublastX/HubLast-Hub/internal/models"
	schemas "github.com/HublastX/HubLast-Hub/internal/schemas"
	"github.com/gin-gonic/gin"
)

// CreateProject godoc
// @Summary Create a new project
// @Description Creates a new project with the provided details
// @Tags projects
// @Accept json
// @Produce json
// @Param project body CreateProjectRequest true "Project details"
// @Success 201 {object} map[string]interface{} "Returns message and created project"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/projects [post]
func (c *ProjectController) CreateProject(ctx *gin.Context) {
	var req schemas.CreateProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
		return
	}

	isAdmin := role == models.AdminRole

	project := &models.Project{
		Title:         req.Title,
		Description:   req.Description,
		FrontendTech:  req.FrontendTech,
		BackendTech:   req.BackendTech,
		EstimatedTime: req.EstimatedTime,
		DeliveryDate:  req.DeliveryDate,
	}

	err := c.projectService.CreateProject(project, userID.(uint), isAdmin)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"project": project,
	})
}
