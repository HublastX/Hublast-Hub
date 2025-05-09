package services

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
	"github.com/HublastX/HubLast-Hub/pkg/database"
	"gorm.io/gorm"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
	userRepo    *repository.UserRepository
	db          *gorm.DB
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		projectRepo: repository.NewProjectRepository(),
		userRepo:    repository.NewUserRepository(),
		db:          database.GetDB(),
	}
}

func (s *ProjectService) CreateProject(project *models.Project, creatorID uint, isAdmin bool) error {

	if !isAdmin {
		project.Status = models.Pending
	} else {
		project.Status = models.Approved
	}

	project.ResponsibleUserID = creatorID

	err := s.projectRepo.Create(project)
	if err != nil {
		return err
	}

	return s.projectRepo.AddUserToProject(project.ID, creatorID)
}

func (s *ProjectService) GetProjectByID(id uint) (*models.Project, error) {
	return s.projectRepo.FindByID(id)
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	return s.projectRepo.FindAll()
}

func (s *ProjectService) GetProjectsByStatus(status models.ProjectStatus) ([]models.Project, error) {
	return s.projectRepo.FindByStatus(status)
}

func (s *ProjectService) UpdateProject(project *models.Project) error {
	return s.projectRepo.Update(project)
}

func (s *ProjectService) DeleteProject(id uint) error {
	return s.projectRepo.Delete(id)
}

func (s *ProjectService) ApproveProject(projectID uint) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	project.Status = models.Approved
	return s.projectRepo.Update(project)
}

func (s *ProjectService) RejectProject(projectID uint) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	project.Status = models.Rejected
	return s.projectRepo.Update(project)
}

func (s *ProjectService) AddUserToProject(projectID, userID uint) error {

	_, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return errors.New("project not found")
	}

	_, err = s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	return s.projectRepo.AddUserToProject(projectID, userID)
}

func (s *ProjectService) RemoveUserFromProject(projectID, userID uint) error {
	return s.projectRepo.RemoveUserFromProject(projectID, userID)
}

func (s *ProjectService) GetUserProjects(userID uint) ([]models.Project, error) {
	return s.projectRepo.FindProjectsByUserID(userID)
}

func (s *ProjectService) SetProjectResponsible(projectID, userID uint) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	project.ResponsibleUserID = userID
	return s.projectRepo.Update(project)
}

func (s *ProjectService) AddFrontendTechToProject(projectID uint, techName string) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	var tech models.Technology
	result := s.db.Where("name = ? AND type = ?", techName, models.FrontendTech).First(&tech)
	if result.Error != nil {
		return result.Error
	}

	return s.db.Model(project).Association("Technologies").Append(&tech)
}

func (s *ProjectService) AddBackendTechToProject(projectID uint, techName string) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	var tech models.Technology
	result := s.db.Where("name = ? AND type = ?", techName, models.BackendTech).First(&tech)
	if result.Error != nil {
		return result.Error
	}

	return s.db.Model(project).Association("Technologies").Append(&tech)
}

func (s *ProjectService) GetFrontendTechsForProject(projectID uint) ([]models.Technology, error) {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	var techs []models.Technology
	err = s.db.Model(project).Association("Technologies").Find(&techs, "type = ?", models.FrontendTech)
	return techs, err
}

func (s *ProjectService) GetBackendTechsForProject(projectID uint) ([]models.Technology, error) {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return nil, err
	}

	var techs []models.Technology
	err = s.db.Model(project).Association("Technologies").Find(&techs, "type = ?", models.BackendTech)
	return techs, err
}

func (s *ProjectService) ClearProjectTechnologies(projectID uint) error {
	project, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return err
	}

	return s.db.Model(project).Association("Technologies").Clear()
}
