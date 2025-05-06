package projectcontrollers

import "github.com/HublastX/HubLast-Hub/internal/services"

type ProjectController struct {
	projectService *services.ProjectService
}

func NewProjectController() *ProjectController {
	return &ProjectController{
		projectService: services.NewProjectService(),
	}
}
