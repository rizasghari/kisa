package kisa

import (
	"fmt"
	"kisa-url-shortner/configs"
	"kisa-url-shortner/internal/handlers"
	"kisa-url-shortner/internal/servers"
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
	cfg := configs.GetConfig()
	db := servers.GetDB(cfg)
	http := servers.NewHttpServer(cfg, db, handlers.NewUserHandler(), handlers.NewHtmlHandler())
	http.Start()
}
