package services

import (
	"kisa-url-shortner/internal/models"
	"kisa-url-shortner/internal/repositories"
	"kisa-url-shortner/internal/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(user *models.User) error {
	hashed, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashed
	return us.userRepository.CreateUser(user)
}

func (us *UserService) Login(email, password string) (*models.User, error) {
	return us.userRepository.Login(email, password)
}
