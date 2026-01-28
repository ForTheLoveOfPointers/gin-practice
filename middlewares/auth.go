package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimSpace(ctx.GetHeader("Authorization"))

		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "missing auth token"})
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte("MYSECRET"), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			ctx.AbortWithError(500, err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["name"], claims["email"])
		} else {
			fmt.Println(err)
		}

		ctx.Next()
	}
}
