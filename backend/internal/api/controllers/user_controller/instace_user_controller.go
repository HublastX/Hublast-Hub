package usercontrollers

import "github.com/HublastX/HubLast-Hub/internal/services"

type UserController struct {
	authService *services.AuthService
	userService *services.UserService
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
	return &UserController{
		authService: services.NewAuthService(),
		userService: services.NewUserService(),
	}
}
