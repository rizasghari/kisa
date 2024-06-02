package controllers

import (
	"github.com/gin-gonic/gin"
	"kisa-url-shortner/internal/models"
	"kisa-url-shortner/internal/services"
	"kisa-url-shortner/internal/utils"
	web "kisa-url-shortner/web/templ"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("eycEW3OKV+axBFZQL4cpbAVRFMhSEc+xRrcHXxhTM8U=")

type Controller struct {
	userService *services.UserService
}

func NewController(userService *services.UserService) *Controller {
	return &Controller{
		userService: userService,
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

	user, err := c.userService.Login(email, password)
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
	err := c.userService.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	ctx.Redirect(http.StatusFound, "/")
}
