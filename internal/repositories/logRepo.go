package repositories

import (
	"gorm.io/gorm"
	"kisa/internal/models"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (lr *LogRepository) CreateLog(log *models.Log) error {
	return lr.db.Create(log).Error
}
