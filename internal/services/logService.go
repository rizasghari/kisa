package services

import (
	"kisa/internal/models"
	"kisa/internal/repositories"
)

type LogService struct {
	logRepository *repositories.LogRepository
}

func NewLogService(logRepository *repositories.LogRepository) *LogService {
	return &LogService{
		logRepository: logRepository,
	}
}

func (ls *LogService) CreateLog(log *models.Log) error {
	return ls.logRepository.CreateLog(log)
}
