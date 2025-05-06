package roadmapcontrollers

import "github.com/HublastX/HubLast-Hub/internal/services"

type RoadmapController struct {
	roadmapService *services.RoadmapService
}

func NewRoadmapController() *RoadmapController {
	return &RoadmapController{
		roadmapService: services.NewRoadmapService(),
	}
}
