package routers

import (
	"testpkg/ginserver/controller"
	"testpkg/ginserver/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
	userService     service.UserService        = service.NewUser()
	userController  controller.UserController  = controller.NewUser(userService)
)

func SetupPublicRouters(rg *gin.RouterGroup) {
	RegisterUsersRouter(rg, &userController)
}

// Refactor this afterwards, or it will become a nightmare for more controllers
func SetupPrivateRouters(rg *gin.RouterGroup) {
	RegisterVideosRouter(rg, &videoController)
}
