package routers

import (
	"testpkg/ginserver/controller"

	"github.com/gin-gonic/gin"
)

// Refactor this afterwards, or it will become a nightmare for more controllers
func SetupRouters(rg *gin.RouterGroup, videoController *controller.VideoController) {
	RegisterVideosRouter(rg, videoController)
}
