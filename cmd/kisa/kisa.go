package kisa

import (
	"fmt"
	"kisa/configs"
	"kisa/internal/cli"
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
	logRepository := repositories.NewLogRepository(db)

	authenticationService := services.NewAuthenticationService(userRepository)
	shortenerService := services.NewShortenerService(urlRepository)
	logService := services.NewLogService(logRepository)
	redisService := services.InitializeRedisClient(config)

	cliChan := make(chan bool)
	_cli := cli.NewCli(shortenerService)
	go _cli.Run(cliChan)

	startHTTPServer := <-cliChan

	if startHTTPServer {
		controller := controllers.NewController(
			authenticationService,
			shortenerService,
			logService,
			redisService,
		)
		httpServer := http.NewHttpServer(
			config,
			db,
			controller,
		)
		httpServer.Start()
	} else {
		fmt.Println("Kisa lost it's way!")
	}
}
