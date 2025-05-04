package repository

import (
	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/pkg/database"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{DB: database.DB}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) FindByID(id uint) (*models.Project, error) {
	var project models.Project
	err := r.DB.Preload("ResponsibleUser").Preload("Users").First(&project, id).Error
	return &project, err
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return r.DB.Save(project).Error
}

func (r *ProjectRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Project{}, id).Error
}

func (r *ProjectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Preload("ResponsibleUser").Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) FindByStatus(status models.ProjectStatus) ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Preload("ResponsibleUser").Where("status = ?", status).Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) AddUserToProject(projectID, userID uint) error {
	return r.DB.Exec("INSERT INTO user_projects (project_id, user_id) VALUES (?, ?)", projectID, userID).Error
}

func (r *ProjectRepository) RemoveUserFromProject(projectID, userID uint) error {
	return r.DB.Exec("DELETE FROM user_projects WHERE project_id = ? AND user_id = ?", projectID, userID).Error
}

func (r *ProjectRepository) FindProjectsByUserID(userID uint) ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("user_projects.user_id = ?", userID).
		Preload("ResponsibleUser").
		Find(&projects).Error
	return projects, err
}
