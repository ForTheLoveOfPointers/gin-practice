package routers

import (
	"testpkg/ginserver/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRouter(rg *gin.RouterGroup, userController *controller.UserController) {
	users := rg.Group("/users")

	users.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(200, (*userController).Register(ctx))
	})

	users.POST("/login", func(ctx *gin.Context) {
		token, err := (*userController).Login(ctx)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, gin.H{"token": token})
	})
}
