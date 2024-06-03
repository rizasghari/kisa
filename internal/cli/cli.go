package cli

import (
	"flag"
	"fmt"
	"kisa/internal/models"
	"kisa/internal/services"
	"kisa/internal/utils"
)

type Cli struct {
	shortenerService *services.ShortenerService
}

func NewCli(shortenerService *services.ShortenerService) *Cli {
	return &Cli{
		shortenerService: shortenerService,
	}
}

func (c *Cli) Run(startHTTPServerCh chan bool) {
	var shorten bool
	var starHttpServer bool

	flag.BoolVar(&shorten, "cli-mode", false, "Shorten url")
	flag.BoolVar(&starHttpServer, "http-mode", false, "Also start http server")

	flag.Parse()

	if shorten && starHttpServer {
		fmt.Println("Cannot shorten and start http server at the same time")
		startHTTPServerCh <- starHttpServer
		return
	}

	if shorten {
		var originalUrl string
		fmt.Print("Enter url: ")
		_, err := fmt.Scanln(&originalUrl)
		if err != nil {
			startHTTPServerCh <- starHttpServer
			return
		}

		err = utils.ValidateUrl(originalUrl)
		if err != nil {
			startHTTPServerCh <- starHttpServer
			return
		}

		url := &models.URL{}
		url.OriginalURL = originalUrl
		shortURL := c.shortenerService.GenerateShortURL(url.OriginalURL)
		url.ShortURL = shortURL
		url.UserID = "cli"

		_, err = c.shortenerService.Shorten(url)
		if err != nil {
			startHTTPServerCh <- starHttpServer
			return
		}
		fmt.Printf("Short url: %s\n", utils.GetFullShortURL(url.ShortURL))
		startHTTPServerCh <- starHttpServer
	}

	if starHttpServer {
		startHTTPServerCh <- starHttpServer
	}

	return
}
