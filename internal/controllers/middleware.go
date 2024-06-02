package controllers

import (
	"github.com/gin-gonic/gin"
	"kisa-url-shortner/internal/utils"
	"log"
	"net/http"
	"strings"
)

func (c *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		log.Println("token: ", tokenString)

		if tokenString == "" {
			log.Println("token not found")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		if strings.Contains(tokenString, "Bearer") {
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		}

		claims, err := utils.VerifyToken(tokenString, jwtKey)
		if err != nil {
			log.Println("invalid token")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		log.Println("valid token")

		ctx.Set("user_id", claims.ID)
		ctx.Set("user_email", claims.Email)
		ctx.Next()
	}
}
