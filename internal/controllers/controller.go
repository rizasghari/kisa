package controllers

import (
	"github.com/gin-gonic/gin"
	"kisa/internal/models"
	"kisa/internal/services"
	"kisa/internal/utils"
	web "kisa/web/templ"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("eycEW3OKV+axBFZQL4cpbAVRFMhSEc+xRrcHXxhTM8U=")

type Controller struct {
	authenticationService *services.AuthenticationService
	shortenerService      *services.ShortenerService
	logService            *services.LogService
}

func NewController(
	authenticationService *services.AuthenticationService,
	shortenerService *services.ShortenerService,
	logService *services.LogService,
) *Controller {
	return &Controller{
		authenticationService: authenticationService,
		shortenerService:      shortenerService,
		logService:            logService,
	}
}

func (c *Controller) GetIndexPage(ctx *gin.Context) {
	email := ctx.GetString("user_email")
	if email != "" {
		ctx.HTML(http.StatusOK, "", web.Index(true, "home"))

	} else {
		ctx.HTML(http.StatusOK, "", web.Index(false, "home"))
	}
}

func (c *Controller) NotFound(ctx *gin.Context) {
	if utils.IsAuthenticated(ctx, jwtKey) {
		log.Println("NotFound - user already authenticated")
		ctx.HTML(http.StatusNotFound, "", web.Index(true, "404"))
	} else {
		log.Println("NotFound - user not authenticated")
		ctx.HTML(http.StatusNotFound, "", web.Index(false, "404"))
	}
}

func (c *Controller) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", web.Index(false, "login"))
}

func (c *Controller) SignupPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", web.Index(false, "signup"))
}

func (c *Controller) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	user, err := c.authenticationService.Login(email, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiration := time.Now().Add(time.Hour * 24).Unix()
	token, err := utils.CreateJwtToken(user.ID, user.Email, jwtKey, time.Unix(expiration, 0))

	ctx.SetCookie("jwt_token", token, int(expiration), "/", "", true, true)
	ctx.HTML(http.StatusOK, "", web.Index(true, "home"))
}

func (c *Controller) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt_token", "", -1, "/", "", true, true)
	ctx.HTML(http.StatusOK, "", web.Index(false, "login"))
}

func (c *Controller) Signup(ctx *gin.Context) {
	var user models.User
	user.Email = ctx.PostForm("email")
	user.PasswordHash = ctx.PostForm("password")
	err := c.authenticationService.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Redirect(http.StatusFound, "/")
}

func (c *Controller) Shorten(ctx *gin.Context) {
	var url models.URL
	url.OriginalURL = ctx.PostForm("url")
	url.UserID = ctx.GetString("user_id")
	short, err := c.shortenerService.Shorten(&url)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.HTML(http.StatusOK, "", web.Result(utils.GetFullShortURL(short)))
}

func (c *Controller) RedirectToOriginalURL(ctx *gin.Context) {
	shortURL := ctx.Param("short")
	URL, err := c.shortenerService.GetOriginalURL(shortURL)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	urlLog := models.Log{
		UrlID:     URL.ID,
		Referrer:  ctx.Request.Referer(),
		UserAgent: ctx.Request.UserAgent(),
		IP:        utils.GetClientIpAddr(ctx.Request),
	}
	err = c.logService.CreateLog(&urlLog)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusFound, URL.OriginalURL)
}
