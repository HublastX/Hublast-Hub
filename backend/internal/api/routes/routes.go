package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	projectcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/project_controller"
	roadmapcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/roadmap_controller"
	usercontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/user_controller"
	"github.com/HublastX/HubLast-Hub/internal/api/middleware"
	"github.com/HublastX/HubLast-Hub/internal/services"
	"github.com/HublastX/HubLast-Hub/pkg/database"
)

func SetupRoutes(router *gin.Engine) {

	router.Use(middleware.CORSMiddleware())

	projectService := services.NewProjectService()
	db := database.DB
	projectController := projectcontrollers.NewProjectController(*projectService, db)
	roadmapController := roadmapcontrollers.NewRoadmapController()
	userController := usercontrollers.NewUserController()

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	api := router.Group("/api")

	SetupUserRoutes(api, userController)
	SetupProjectRoutes(api, projectController)
	SetupRoadmapRoutes(api, roadmapController)
}
