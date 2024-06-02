package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"kisa-url-shortner/internal/models"
	"kisa-url-shortner/internal/services"
	web "kisa-url-shortner/web/templ"
	"net/http"
	"time"
)

var jwtKey = []byte("eycEW3OKV+axBFZQL4cpbAVRFMhSEc+xRrcHXxhTM8U=")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
type Controller struct {
	userService *services.UserService
}

func NewController(userService *services.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) GetIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", web.Index(true, "home"))
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

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	ctx.SetCookie("jwt_token", tokenString, int(expirationTime.Sub(time.Now()).Seconds()), "/", "",
		false, true)
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
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
