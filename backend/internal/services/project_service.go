package services

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
	userRepo    *repository.UserRepository
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		projectRepo: repository.NewProjectRepository(),
		userRepo:    repository.NewUserRepository(),
	}
}

func (s *ProjectService) CreateProject(project *models.Project, creatorID uint, isAdmin bool) error {
	// If creator is not admin, set status to pending
	if !isAdmin {
		project.Status = models.Pending
	} else {
		project.Status = models.Approved
	}

	// Set responsible user
	project.ResponsibleUserID = creatorID

	// Create project
	err := s.projectRepo.Create(project)
	if err != nil {
		return err
	}

	// Add creator to project users
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
	// Check if project exists
	_, err := s.projectRepo.FindByID(projectID)
	if err != nil {
		return errors.New("project not found")
	}

	// Check if user exists
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
