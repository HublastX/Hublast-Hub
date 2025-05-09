package usercontrollers

import "github.com/HublastX/HubLast-Hub/internal/services"

type UserController struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		authService: services.NewAuthService(),
		userService: services.NewUserService(),
	}
}
