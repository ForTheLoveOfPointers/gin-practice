package main

import (
	"fmt"
	"io"
	"os"
	"testpkg/ginserver/controller"
	"testpkg/ginserver/middlewares"
	"testpkg/ginserver/routers"
	"testpkg/ginserver/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	file, err := os.Create("gin.log")

	if err != nil {
		errFmt := fmt.Errorf("ERROR || Unable to create or access gin.log file --- %w", err)
		fmt.Println(errFmt)
		return
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	protected := server.Group("/my-account")
	protected.Use(middlewares.Auth())
	{
		routers.SetupRouters(protected, &videoController)
	}

	server.Run(":3000")
}
