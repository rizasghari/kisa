package servers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kisa-url-shortner/internal/models"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=user password=password dbname=url_shortener port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.URL{}, &models.Log{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
