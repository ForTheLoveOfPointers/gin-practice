package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.TrimSpace(ctx.GetHeader("Authorization"))

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "missing auth token"})
			return
		}

		/**
		Validate token
		*/

		ctx.Next()
	}
}
