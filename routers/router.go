package routers

import (
	"testpkg/ginserver/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(rg *gin.RouterGroup, videoController *controller.VideoController) {
	RegisterVideosRouter(rg, videoController)
}
