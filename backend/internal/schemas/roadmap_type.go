package schemas

import "github.com/HublastX/HubLast-Hub/internal/models"

type CreateRoadmapRequest struct {
	Title       string                 `json:"title" binding:"required"`
	Area        models.RoadmapArea     `json:"area" binding:"required"`
	Difficulty  models.DifficultyLevel `json:"difficulty" binding:"required"`
	Content     string                 `json:"content" binding:"required"`
	CourseLinks string                 `json:"course_links"`
}
