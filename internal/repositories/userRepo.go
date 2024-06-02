package repositories

import (
	"gorm.io/gorm"
	"kisa-url-shortner/internal/models"
	"kisa-url-shortner/internal/utils"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetUser(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Login(email, password string) (*models.User, error) {
	user, err := ur.GetUser(email)
	if err != nil {
		return nil, err
	}
	if err := utils.CompareHashAndPassword(user.PasswordHash, password); err != nil {
		return nil, err
	}
	return user, nil
}
