package main

import (
	"fmt"
	"io"
	"os"
	"testpkg/ginserver/controller"
	"testpkg/ginserver/middlewares"
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
		errFmt := fmt.Errorf("ERROR || Unable to create gin.log file --- %w", err)
		fmt.Println(errFmt)
		return
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":3000")
}
