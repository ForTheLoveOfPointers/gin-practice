package controller

import (
	"testpkg/ginserver/entity"
	"testpkg/ginserver/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(*gin.Context) (entity.User, error)
	Login(*gin.Context) (string, error)
}

type userController struct {
	service service.UserService
}

func NewUser(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Register(ctx *gin.Context) (entity.User, error) {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.Register(user)
}

func (c *userController) Login(ctx *gin.Context) (string, error) {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.Login(user)
}
