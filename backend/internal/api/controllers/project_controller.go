package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/services"
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	projectService *services.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		projectService: services.NewProjectService(),
	}
}

// Add these struct definitions after CreateProjectRequest
type UserIDRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

type CreateProjectRequest struct {
	Title         string    `json:"title" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	FrontendTech  string    `json:"frontend_tech" binding:"required"`
	BackendTech   string    `json:"backend_tech" binding:"required"`
	EstimatedTime int       `json:"estimated_time" binding:"required"`
	DeliveryDate  time.Time `json:"delivery_date" binding:"required"`
}

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
	var req CreateProjectRequest
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

// GetProject godoc
// @Summary Get a specific project
// @Description Retrieves a project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{} "Returns project details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Project not found"
// @Security ApiKeyAuth
// @Router /api/projects/{id} [get]
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

// GetAllProjects godoc
// @Summary Get all projects
// @Description Retrieves all projects, optionally filtered by status
// @Tags projects
// @Accept json
// @Produce json
// @Param status query string false "Filter projects by status"
// @Success 200 {object} map[string]interface{} "Returns list of projects"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/projects [get]
func (c *ProjectController) GetAllProjects(ctx *gin.Context) {
	status := ctx.Query("status")

	var projects []models.Project
	var err error

	if status != "" {
		projects, err = c.projectService.GetProjectsByStatus(models.ProjectStatus(status))
	} else {
		projects, err = c.projectService.GetAllProjects()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}

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

	var req CreateProjectRequest
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

// DeleteProject godoc
// @Summary Delete a project
// @Description Deletes a project by its ID (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Project deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id} [delete]
func (c *ProjectController) DeleteProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete projects"})
		return
	}

	err = c.projectService.DeleteProject(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// ApproveProject godoc
// @Summary Approve a project
// @Description Approves a pending project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Project approved successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/approve [post]
func (c *ProjectController) ApproveProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can approve projects"})
		return
	}

	err = c.projectService.ApproveProject(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project approved successfully"})
}

// RejectProject godoc
// @Summary Reject a project
// @Description Rejects a pending project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Project rejected successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/reject [post]
func (c *ProjectController) RejectProject(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can reject projects"})
		return
	}

	err = c.projectService.RejectProject(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project rejected successfully"})
}

// AddUserToProject godoc
// @Summary Add user to project
// @Description Adds a user to a specific project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param user body UserIDRequest true "User ID to add"
// @Success 200 {object} map[string]string "User added to project successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/users [post]
func (c *ProjectController) AddUserToProject(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can add users to projects"})
		return
	}

	err = c.projectService.AddUserToProject(uint(projectID), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User added to project successfully"})
}

// RemoveUserFromProject godoc
// @Summary Remove user from project
// @Description Removes a user from a specific project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param user body UserIDRequest true "User ID to remove"
// @Success 200 {object} map[string]string "User removed from project successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/users [delete]
func (c *ProjectController) RemoveUserFromProject(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can remove users from projects"})
		return
	}

	err = c.projectService.RemoveUserFromProject(uint(projectID), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User removed from project successfully"})
}

// GetProjectUsers godoc
// @Summary Get users in a project
// @Description Retrieves all users associated with a specific project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]interface{} "Returns project details with users"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Project not found"
// @Security ApiKeyAuth
// @Router /api/projects/{id}/users [get]
func (c *ProjectController) GetProjectUsers(ctx *gin.Context) {
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

	var users []gin.H
	for _, user := range project.Users {
		users = append(users, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"project_id":    project.ID,
		"project_title": project.Title,
		"responsible_user": gin.H{
			"id":       project.ResponsibleUser.ID,
			"username": project.ResponsibleUser.Username,
		},
		"users": users,
	})
}

// SetProjectResponsible godoc
// @Summary Set project responsible
// @Description Sets a user as responsible for a project (admin only)
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param user body UserIDRequest true "User ID to set as responsible"
// @Success 200 {object} map[string]string "Project responsible set successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/admin/projects/{id}/responsible [post]
func (c *ProjectController) SetProjectResponsible(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can set project responsible"})
		return
	}

	err = c.projectService.SetProjectResponsible(uint(projectID), req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Project responsible set successfully"})
}

// RequestJoinProject godoc
// @Summary Request to join a project
// @Description Allows the current user to request joining a specific project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Join request successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/projects/{id}/join [post]
func (c *ProjectController) RequestJoinProject(ctx *gin.Context) {
	projectID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// In a real implementation, you might create a join request record
	// For simplicity, we'll just add the user to the project directly
	err = c.projectService.AddUserToProject(uint(projectID), userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Join request submitted successfully"})
}

// GetUserProjects godoc
// @Summary Get current user's projects
// @Description Retrieves all projects associated with the authenticated user
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns list of user's projects"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security ApiKeyAuth
// @Router /api/user/projects [get]
func (c *ProjectController) GetUserProjects(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projects, err := c.projectService.GetUserProjects(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"projects": projects})
}
