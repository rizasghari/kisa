package kisa

import (
	"fmt"
	"kisa/configs"
	"kisa/internal/controllers"
	"kisa/internal/repositories"
	"kisa/internal/servers"
	"kisa/internal/servers/http"
	"kisa/internal/services"
	"sync"
)

var (
	kisa *Kisa
	once sync.Once
)

type Kisa struct{}

func App() *Kisa {
	once.Do(func() {
		kisa = &Kisa{}
	})
	return kisa
}

func (k *Kisa) LetsGo() {
	fmt.Println("App is running")

	config := configs.GetConfig()

	db := servers.GetDB(config)

	userRepository := repositories.NewUserRepository(db)
	urlRepository := repositories.NewUrlRepository(db)

	authenticationService := services.NewAuthenticationService(userRepository)
	shortenerService := services.NewShortenerService(urlRepository)

	httpServer := http.NewHttpServer(config, db, controllers.NewController(authenticationService, shortenerService))
	httpServer.Start()
}
