package routes

import (
	projectcontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/project_controller"
	"github.com/HublastX/HubLast-Hub/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(router *gin.RouterGroup, projectController *projectcontrollers.ProjectController) {

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/projects", projectController.CreateProject)
		protected.GET("/projects", projectController.GetAllProjects)
		protected.GET("/projects/:id", projectController.GetProject)
		protected.GET("/user/projects", projectController.GetUserProjects)
		protected.POST("/projects/:id/join", projectController.RequestJoinProject)
		protected.GET("/projects/:id/users", projectController.GetProjectUsers)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.PUT("/projects/:id", projectController.UpdateProject)
		admin.DELETE("/projects/:id", projectController.DeleteProject)
		admin.POST("/projects/:id/approve", projectController.ApproveProject)
		admin.POST("/projects/:id/reject", projectController.RejectProject)
		admin.POST("/projects/:id/users", projectController.AddUserToProject)
		admin.DELETE("/projects/:id/users", projectController.RemoveUserFromProject)
		admin.POST("/projects/:id/responsible", projectController.SetProjectResponsible)
	}
}
