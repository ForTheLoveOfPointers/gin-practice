package routers

import (
	"testpkg/ginserver/controller"
	"testpkg/ginserver/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
	userService     service.UserService
	userController  controller.UserController
)

func SetupPublicRouters(rg *gin.RouterGroup, db *gorm.DB) {
	userService = service.NewUser(db)
	userController = controller.NewUser(userService)
	RegisterUsersRouter(rg, &userController)
}

func SetupPrivateRouters(rg *gin.RouterGroup) {
	RegisterVideosRouter(rg, &videoController)
}
