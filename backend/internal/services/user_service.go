package services

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
	"github.com/HublastX/HubLast-Hub/pkg/utils"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) ChangePassword(userID uint, currentPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(currentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := utils.GenerateHashPassword(newPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.Update(user)
}

func (s *UserService) PromoteToAdmin(userID uint) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	user.Role = models.AdminRole
	return s.userRepo.Update(user)
}
