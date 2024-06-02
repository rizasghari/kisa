package controllers

import (
	"github.com/gin-gonic/gin"
	"kisa-url-shortner/internal/utils"
	"log"
	"net/http"
	"strings"
)

func (c *Controller) AlreadyAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := ctx.Cookie("jwt_token")
		if err != nil {
			log.Println("JWT token not provided")
			ctx.Next()
		}
		if _, err = utils.VerifyToken(jwtToken, jwtKey); err == nil {
			log.Println("user already authenticated")
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		ctx.Next()
	}
}

func (c *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var jwtToken string
		jwtTokenFromHeader := ctx.GetHeader("Authorization")
		if jwtTokenFromHeader != "" {
			if strings.Contains(jwtTokenFromHeader, "Bearer") {
				jwtToken = strings.Replace(jwtTokenFromHeader, "Bearer ", "", 1)
			} else {
				jwtToken = jwtTokenFromHeader
			}
		} else {
			jwtTokenFromCookie, err := ctx.Cookie("jwt_token")
			if err != nil {
				ctx.Redirect(http.StatusFound, "/login")
				return
			}
			jwtToken = jwtTokenFromCookie
		}

		if jwtToken == "" {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		claims, err := utils.VerifyToken(jwtToken, jwtKey)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		ctx.Set("user_id", claims.ID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("authenticated", true)
		ctx.Next()
	}
}
