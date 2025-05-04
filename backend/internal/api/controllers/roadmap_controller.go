package controllers

import (
	"net/http"
	"strconv"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/services"
	"github.com/gin-gonic/gin"
)

type RoadmapController struct {
	roadmapService *services.RoadmapService
}

func NewRoadmapController() *RoadmapController {
	return &RoadmapController{
		roadmapService: services.NewRoadmapService(),
	}
}

type CreateRoadmapRequest struct {
	Title       string                 `json:"title" binding:"required"`
	Area        models.RoadmapArea     `json:"area" binding:"required"`
	Difficulty  models.DifficultyLevel `json:"difficulty" binding:"required"`
	Content     string                 `json:"content" binding:"required"`
	CourseLinks string                 `json:"course_links"`
}

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
	var req CreateRoadmapRequest
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

// UpdateRoadmap godoc
// @Summary Update roadmap information
// @Description Update an existing roadmap (admin only)
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param id path int true "Roadmap ID"
// @Param roadmap body CreateRoadmapRequest true "Updated roadmap information"
// @Success 200 {object} map[string]interface{} "Roadmap updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - Admin only"
// @Failure 404 {object} map[string]string "Roadmap not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/roadmaps/{id} [put]
func (c *RoadmapController) UpdateRoadmap(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roadmap ID"})
		return
	}

	var req CreateRoadmapRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update roadmaps"})
		return
	}

	roadmap, err := c.roadmapService.GetRoadmapByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Roadmap not found"})
		return
	}

	roadmap.Title = req.Title
	roadmap.Area = req.Area
	roadmap.Difficulty = req.Difficulty
	roadmap.Content = req.Content
	roadmap.CourseLinks = req.CourseLinks

	err = c.roadmapService.UpdateRoadmap(roadmap)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Roadmap updated successfully",
		"roadmap": roadmap,
	})
}

// DeleteRoadmap godoc
// @Summary Delete roadmap
// @Description Delete an existing roadmap (admin only)
// @Tags roadmaps
// @Accept json
// @Produce json
// @Param id path int true "Roadmap ID"
// @Success 200 {object} map[string]string "Roadmap deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 403 {object} map[string]string "Forbidden - Admin only"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /admin/roadmaps/{id} [delete]
func (c *RoadmapController) DeleteRoadmap(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid roadmap ID"})
		return
	}

	role, exists := ctx.Get("role")
	if !exists || role != models.AdminRole {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete roadmaps"})
		return
	}

	err = c.roadmapService.DeleteRoadmap(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Roadmap deleted successfully"})
}
