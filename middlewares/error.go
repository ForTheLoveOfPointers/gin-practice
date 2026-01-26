package middlewares

import (
	"errors"
	servererrors "testpkg/ginserver/server_errors"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err

		var reqErr *servererrors.RequestError
		if errors.As(err, &reqErr) {
			ctx.JSON(reqErr.Code, gin.H{
				"error": reqErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error": "internal server error",
		})
	}
}
