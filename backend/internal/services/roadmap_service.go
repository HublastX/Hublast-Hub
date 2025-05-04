package services

import (
	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
)

type RoadmapService struct {
	roadmapRepo *repository.RoadmapRepository
}

func NewRoadmapService() *RoadmapService {
	return &RoadmapService{
		roadmapRepo: repository.NewRoadmapRepository(),
	}
}

func (s *RoadmapService) CreateRoadmap(roadmap *models.Roadmap) error {
	return s.roadmapRepo.Create(roadmap)
}

func (s *RoadmapService) GetRoadmapByID(id uint) (*models.Roadmap, error) {
	return s.roadmapRepo.FindByID(id)
}

func (s *RoadmapService) GetAllRoadmaps() ([]models.Roadmap, error) {
	return s.roadmapRepo.FindAll()
}

func (s *RoadmapService) GetRoadmapsByArea(area models.RoadmapArea) ([]models.Roadmap, error) {
	return s.roadmapRepo.FindByArea(area)
}

func (s *RoadmapService) UpdateRoadmap(roadmap *models.Roadmap) error {
	return s.roadmapRepo.Update(roadmap)
}

func (s *RoadmapService) DeleteRoadmap(id uint) error {
	return s.roadmapRepo.Delete(id)
}
