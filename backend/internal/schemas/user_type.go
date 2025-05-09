package schemas

import (
	"github.com/HublastX/HubLast-Hub/internal/models"
)

type RegisterRequest struct {
	Username   string            `json:"username" binding:"required" example:"johndoe"`
	Email      string            `json:"email" binding:"required,email" example:"john@example.com"`
	Password   string            `json:"password" binding:"required,min=6" example:"password123"`
	Level      models.Level      `json:"level" binding:"required"  example:"junior"`
	Experience models.Experience `json:"experience" binding:"required"  example:"basic"`
	Employment models.Employment `json:"employment" binding:"required"  example:"backend"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required" example:"password123"`
	NewPassword     string `json:"new_password" binding:"required,min=6" example:"newpassword123"`
}

type UserBasicInfo struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email,omitempty" example:"john@example.com"`
	Role     string `json:"role,omitempty" example:"user"`
}
