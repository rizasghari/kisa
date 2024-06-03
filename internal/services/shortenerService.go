package services

import (
	"crypto/md5"
	"encoding/hex"
	"kisa/internal/models"
	"kisa/internal/repositories"
	"kisa/internal/utils"
)

type ShortenerService struct {
	urlRepository *repositories.UrlRepository
}

func NewShortenerService(urlRepository *repositories.UrlRepository) *ShortenerService {
	return &ShortenerService{
		urlRepository: urlRepository,
	}
}

func (ss *ShortenerService) Shorten(url *models.URL) (string, error) {
	err := utils.ValidateUrl(url.OriginalURL)
	if err != nil {
		return "", err
	}
	url.ShortURL = ss.GenerateShortURL(url.OriginalURL)
	return ss.urlRepository.AddNew(url)
}

func (ss *ShortenerService) GenerateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:6]
}

func (ss *ShortenerService) GetOriginalURL(shortURL string) (*models.URL, error) {

	return ss.urlRepository.GetOriginalURL(shortURL)
}
