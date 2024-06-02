package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kisa-url-shortner/configs"
	"kisa-url-shortner/internal/controllers"
	"kisa-url-shortner/web/templ"
	"log"
)

type Server struct {
	Config     *configs.Config
	Router     *gin.Engine
	db         *gorm.DB
	controller *controllers.Controller
}

func NewHttpServer(
	config *configs.Config,
	dbHandler *gorm.DB,
	httpHandler *controllers.Controller,
) *Server {
	return &Server{
		Config:     config,
		db:         dbHandler,
		controller: httpHandler,
	}
}

func (hs *Server) Start() {
	hs.Router = gin.Default()
	hs.configureStatics()
	hs.configureTempl()
	hs.setRoutes()
	hs.run()
}

func (hs *Server) configureStatics() {
	hs.Router.Static("/web/static", "./web/static")
}

func (hs *Server) configureTempl() {
	ginHtmlRenderer := hs.Router.HTMLRender
	hs.Router.HTMLRender = &web.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	err := hs.Router.SetTrustedProxies(nil)
	if err != nil {
		return
	}
}

func (hs *Server) setRoutes() {

	hs.Router.POST("/signup", hs.controller.Signup)
	hs.Router.POST("/login", hs.controller.Login)

	alreadyAuthenticated := hs.Router.Group("/")
	alreadyAuthenticated.Use(hs.controller.AlreadyAuthenticated())
	{
		alreadyAuthenticated.GET("/login", hs.controller.LoginPage)
		alreadyAuthenticated.GET("/signup", hs.controller.SignupPage)
	}

	authorized := hs.Router.Group("/")
	authorized.Use(hs.controller.AuthMiddleware())
	{
		authorized.GET("/", hs.controller.GetIndexPage)
	}
}

func (hs *Server) run() {
	port := fmt.Sprintf(":%s", hs.Config.Viper.GetString("http.port"))
	err := hs.Router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
