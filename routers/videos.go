package routers

import (
	"testpkg/ginserver/controller"

	"github.com/gin-gonic/gin"
)

func RegisterVideosRouter(rg *gin.RouterGroup, videoController *controller.VideoController) {

	videos := rg.Group("/videos")

	videos.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, *videoController)
	})

	videos.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, (*videoController).Save(ctx))
	})

}
