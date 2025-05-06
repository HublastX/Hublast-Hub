package schemas

import (
	"time"
)

type UserIDRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

type CreateProjectRequest struct {
	Title         string    `json:"title" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	FrontendTech  string    `json:"frontend_tech" binding:"required"`
	BackendTech   string    `json:"backend_tech" binding:"required"`
	EstimatedTime int       `json:"estimated_time" binding:"required"`
	DeliveryDate  time.Time `json:"delivery_date" binding:"required"`
}
