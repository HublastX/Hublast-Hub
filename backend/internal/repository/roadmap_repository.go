package repository

import (
	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/pkg/database"
	"gorm.io/gorm"
)

type RoadmapRepository struct {
	DB *gorm.DB
}

func NewRoadmapRepository() *RoadmapRepository {
	return &RoadmapRepository{DB: database.DB}
}

func (r *RoadmapRepository) Create(roadmap *models.Roadmap) error {
	return r.DB.Create(roadmap).Error
}

func (r *RoadmapRepository) FindByID(id uint) (*models.Roadmap, error) {
	var roadmap models.Roadmap
	err := r.DB.First(&roadmap, id).Error
	return &roadmap, err
}

func (r *RoadmapRepository) Update(roadmap *models.Roadmap) error {
	return r.DB.Save(roadmap).Error
}

func (r *RoadmapRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Roadmap{}, id).Error
}

func (r *RoadmapRepository) FindAll() ([]models.Roadmap, error) {
	var roadmaps []models.Roadmap
	err := r.DB.Find(&roadmaps).Error
	return roadmaps, err
}

func (r *RoadmapRepository) FindByArea(area models.RoadmapArea) ([]models.Roadmap, error) {
	var roadmaps []models.Roadmap
	err := r.DB.Where("area = ?", area).Find(&roadmaps).Error
	return roadmaps, err
}
