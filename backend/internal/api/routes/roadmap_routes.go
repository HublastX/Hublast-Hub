package routes

import (
	roadmapcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/roadmap_controller"
	"github.com/HublastX/HubLast-Hub/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoadmapRoutes(router *gin.RouterGroup, roadmapController *roadmapcontrollers.RoadmapController) {

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/roadmaps", roadmapController.GetAllRoadmaps)
		protected.GET("/roadmaps/:id", roadmapController.GetRoadmap)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/roadmaps", roadmapController.CreateRoadmap)
		admin.PUT("/roadmaps/:id", roadmapController.UpdateRoadmap)
		admin.DELETE("/roadmaps/:id", roadmapController.DeleteRoadmap)
	}
}
