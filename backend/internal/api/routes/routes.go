package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/HublastX/HubLast-Hub/docs"
	projectcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/project_controller"
	roadmapcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/roadmap_controller"
	usercontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/user_controller"
	"github.com/HublastX/HubLast-Hub/internal/api/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	projectController := projectcontrollers.NewProjectController()
	roadmapController := roadmapcontrollers.NewRoadmapController()
	userController := usercontrollers.NewUserController()

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Public routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", userController.Register)
			auth.POST("/login", userController.Login)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			protected.GET("/users", userController.GetAllUsers)
			protected.GET("/users/:id", userController.GetUser)
			protected.PUT("/users/:id", userController.UpdateUser)
			protected.POST("/users/change-password", userController.ChangePassword)

			// Admin-only user routes
			adminUser := protected.Group("/admin/")
			adminUser.Use(middleware.AdminMiddleware())
			{
				adminUser.DELETE("/users/:id", userController.DeleteUser)
				adminUser.POST("/users/:id/promote", userController.PromoteToAdmin)
			}

			// Project routes
			protected.POST("/projects", projectController.CreateProject)
			protected.GET("/projects", projectController.GetAllProjects)
			protected.GET("/projects/:id", projectController.GetProject)
			protected.GET("/user/projects", projectController.GetUserProjects)
			protected.POST("/projects/:id/join", projectController.RequestJoinProject)
			protected.GET("/projects/:id/users", projectController.GetProjectUsers)

			// Admin-only project routes
			adminProject := protected.Group("/admin/")
			adminProject.Use(middleware.AdminMiddleware())
			{
				adminProject.PUT("/projects/:id", projectController.UpdateProject)
				adminProject.DELETE("/projects/:id", projectController.DeleteProject)
				adminProject.POST("/projects/:id/approve", projectController.ApproveProject)
				adminProject.POST("/projects/:id/reject", projectController.RejectProject)
				adminProject.POST("/projects/:id/users", projectController.AddUserToProject)
				adminProject.DELETE("/projects/:id/users", projectController.RemoveUserFromProject)
				adminProject.POST("/projects/:id/responsible", projectController.SetProjectResponsible)
			}

			// Roadmap routes
			protected.GET("/roadmaps", roadmapController.GetAllRoadmaps)
			protected.GET("/roadmaps/:id", roadmapController.GetRoadmap)

			// Admin-only roadmap routes
			adminRoadmap := protected.Group("/admin/")
			adminRoadmap.Use(middleware.AdminMiddleware())
			{
				adminRoadmap.POST("/roadmaps", roadmapController.CreateRoadmap)
				adminRoadmap.PUT("/roadmaps/:id", roadmapController.UpdateRoadmap)
				adminRoadmap.DELETE("/roadmaps/:id", roadmapController.DeleteRoadmap)
			}
		}
	}
}
