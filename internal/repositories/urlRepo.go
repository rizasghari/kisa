package repositories

import (
	"gorm.io/gorm"
	"kisa/internal/models"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (ur *UrlRepository) AddNew(url *models.URL) (string, error) {
	err := ur.db.Create(url).Error
	if err != nil {
		return "", err
	}
	return url.ShortURL, nil
}

func (ur *UrlRepository) GetOriginalURL(shortURL string) (string, error) {
	var url models.URL
	err := ur.db.Where("short_url = ?", shortURL).First(&url).Error
	if err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}
