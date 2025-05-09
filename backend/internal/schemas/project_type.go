package schemas

import (
	"time"
)

type UserIDRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}
type CreateProjectRequest struct {
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	FrontendTechs  []string  `json:"frontend_techs"`
	BackendTechs   []string  `json:"backend_techs"`
	EstimatedTime  int       `json:"estimated_time" binding:"required"`
	DeliveryDate   time.Time `json:"delivery_date" binding:"required"`
	QuantyMaxUsers int       `json:"quanty_max_users" binding:"required"`
	Level          string    `json:"level"`
}

type ProjectUsersResponse struct {
	ProjectID       uint            `json:"project_id" example:"1"`
	ProjectTitle    string          `json:"project_title" example:"Sistema de E-commerce"`
	ResponsibleUser UserBasicInfo   `json:"responsible_user"`
	Users           []UserBasicInfo `json:"users"`
}
