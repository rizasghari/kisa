package services

import (
	"kisa-url-shortner/internal/models"
	"kisa-url-shortner/internal/repositories"
	"kisa-url-shortner/internal/utils"
)

type AuthenticationService struct {
	userRepository *repositories.UserRepository
}

func NewAuthenticationService(userRepository *repositories.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
	}
}

func (as *AuthenticationService) CreateUser(user *models.User) error {
	hashed, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashed
	return as.userRepository.CreateUser(user)
}

func (as *AuthenticationService) Login(email, password string) (*models.User, error) {
	return as.userRepository.Login(email, password)
}
