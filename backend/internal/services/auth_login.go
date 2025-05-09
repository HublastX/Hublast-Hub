package services

import (
	"errors"

	"github.com/HublastX/HubLast-Hub/internal/models"
	"github.com/HublastX/HubLast-Hub/internal/repository"
	"github.com/HublastX/HubLast-Hub/pkg/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

func (s *AuthService) Register(
	username string,
	email string,
	password string,
	level models.Level,
	experience models.Experience,
	employment models.Employment,
) (*models.User, error) {

	_, err := s.userRepo.FindByUsername(username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	_, err = s.userRepo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.GenerateHashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:   username,
		Email:      email,
		Password:   hashedPassword,
		Level:      level,
		Experience: experience,
		Employment: employment,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (string, *models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}
