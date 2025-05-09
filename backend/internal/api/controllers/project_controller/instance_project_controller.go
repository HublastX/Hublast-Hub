package projectcontrollers

import (
	"github.com/HublastX/HubLast-Hub/internal/services"
	"gorm.io/gorm"
)

// Parte do arquivo de definição do ProjectController
type ProjectController struct {
	projectService services.ProjectService
	db             *gorm.DB
}

func NewProjectController(projectService services.ProjectService, db *gorm.DB) *ProjectController {
	return &ProjectController{
		projectService: projectService,
		db:             db,
	}
}
