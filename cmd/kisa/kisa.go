package kisa

import (
	"fmt"
	"kisa-url-shortner/configs"
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
	db := servers.GetDB(configs.GetConfig())
	fmt.Println(db)
}
