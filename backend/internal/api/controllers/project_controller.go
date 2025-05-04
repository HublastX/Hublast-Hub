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

type CreateProjectRequest struct {
	Title         string    `json:"title" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	FrontendTech  string    `json:"frontend_tech" binding:"required"`
	BackendTech   string    `json:"backend_tech" binding:"required"`
	EstimatedTime int       `json:"estimated_time" binding:"required"` // In days
	DeliveryDate  time.Time `json:"delivery_date" binding:"required"`
}

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
