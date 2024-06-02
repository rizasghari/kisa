package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			log.Println("token not found")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil || !token.Valid {
			log.Println("invalid token")
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		log.Println("valid token")

		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
