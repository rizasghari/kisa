package servers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kisa-url-shortner/configs"
	"kisa-url-shortner/internal/handlers"
	"kisa-url-shortner/web/templ"
	"log"
	"net/http"
)

type HttpServer struct {
	Config      *configs.Config
	Router      *gin.Engine
	db          *gorm.DB
	userHandler *handlers.UserHandler
	httpHandler *handlers.HtmlHandler
}

func NewHttpServer(
	config *configs.Config,
	dbHandler *gorm.DB,
	userHandler *handlers.UserHandler,
	httpHandler *handlers.HtmlHandler,
) *HttpServer {
	return &HttpServer{
		Config:      config,
		db:          dbHandler,
		userHandler: userHandler,
		httpHandler: httpHandler,
	}
}

func (hs *HttpServer) Start() {
	hs.Router = gin.Default()
	hs.configureTempl()
	hs.setRoutes()
	hs.run()
}

func (hs *HttpServer) configureTempl() {
	ginHtmlRenderer := hs.Router.HTMLRender
	hs.Router.HTMLRender = &web.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	err := hs.Router.SetTrustedProxies(nil)
	if err != nil {
		return
	}
}

func (hs *HttpServer) setRoutes() {
	hs.Router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", hs.httpHandler.GetIndexPage())
	})
}

func (hs *HttpServer) run() {
	port := fmt.Sprintf(":%s", hs.Config.Viper.GetString("http.port"))
	err := hs.Router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
