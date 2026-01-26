package routers

import (
	"errors"
	"testpkg/ginserver/controller"
	servererrors "testpkg/ginserver/server_errors"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(rg *gin.RouterGroup, userController *controller.UserController) {
	users := rg.Group("/users")

	users.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(200, (*userController).Register(ctx))
	})

	users.POST("/login", func(ctx *gin.Context) {
		token, err := (*userController).Login(ctx)
		var reqError *servererrors.RequestError
		if err != nil {
			if errors.As(err, &reqError) {
				ctx.JSON(reqError.Code, gin.H{
					"error": reqError.Message,
				})
			}
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, gin.H{"token": token})
	})
}
