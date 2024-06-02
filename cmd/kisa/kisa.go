package kisa

import (
	"fmt"
	"kisa-url-shortner/configs"
	"kisa-url-shortner/internal/controllers"
	"kisa-url-shortner/internal/repositories"
	"kisa-url-shortner/internal/servers"
	"kisa-url-shortner/internal/servers/http"
	"kisa-url-shortner/internal/services"
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

	userService := services.NewUserService(userRepository)

	httpServer := http.NewHttpServer(config, db, controllers.NewController(userService))
	httpServer.Start()
}
