package routes

import (
	usercontrollers "github.com/HublastX/HubLast-Hub/internal/api/controllers/user_controller"
	"github.com/HublastX/HubLast-Hub/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userController *usercontrollers.UserController) {

	auth := router.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetAllUsers)
		protected.GET("/users/:id", userController.GetUser)
		protected.PUT("/users/:id", userController.UpdateUser)
		protected.POST("/users/change-password", userController.ChangePassword)
	}

	adminUser := router.Group("/admin")
	adminUser.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminUser.DELETE("/users/:id", userController.DeleteUser)
		adminUser.POST("/users/:id/promote", userController.PromoteToAdmin)
	}
}
