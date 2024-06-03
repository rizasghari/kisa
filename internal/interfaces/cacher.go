package interfaces

import "kisa/internal/models"

type Cacher interface {
	Get(key string) (*models.URL, error)
	Set(key string, value *models.URL) error
	Delete(key string) error
	Flush() error
	Check(key string) bool
}
