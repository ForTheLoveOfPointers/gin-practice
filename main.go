package main

import (
	"testpkg/ginserver/controller"
	"testpkg/ginserver/middlewares"
	"testpkg/ginserver/routers"
	"testpkg/ginserver/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
	userService     service.UserService        = service.NewUser()
	userController  controller.UserController  = controller.NewUser(userService)
)

func main() {
	middlewares.SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	protected := server.Group("/my-account")
	protected.Use(middlewares.Auth())
	{
		routers.SetupRouters(protected, &videoController, &userController)
	}

	server.Run(":3000")
}
