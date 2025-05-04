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

// Empty method to match existing model
func (c *RoadmapController) CreateTask(ctx *gin.Context) {}

// Empty method to match existing model
func (c *RoadmapController) GetTask(ctx *gin.Context) {}

// Empty method to match existing model
func (c *RoadmapController) UpdateTask(ctx *gin.Context) {}

// Empty method to match existing model
func (c *RoadmapController) DeleteTask(ctx *gin.Context) {}
