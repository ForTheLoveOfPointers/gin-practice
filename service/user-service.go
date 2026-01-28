package service

import (
	"fmt"
	"testpkg/ginserver/entity"
	servererrors "testpkg/ginserver/server_errors"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService interface {
	Register(entity.User) entity.User
	Login(entity.User) (string, error)
}

type userService struct {
	db *gorm.DB
}

func NewUser() *userService {
	return &userService{}
}

func (e *userService) Register(usr entity.User) entity.User {
	e.db.Create(usr)
	return usr
}

func (e *userService) Login(usr entity.User) (string, error) {
	var userEnt entity.User
	result := e.db.First(&userEnt, usr.Id)

	if result.Error != nil {
		err := servererrors.RequestError{Code: 404, Message: "User not in database"}
		return "", &err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":  userEnt.Name,
			"email": userEnt.Email,
		},
	)

	token_string, err := token.SignedString([]byte("MYSECRET"))
	if err != nil {
		return "", &servererrors.RequestError{Code: 500, Message: "Failed at jwt signing"}
	}

	fmt.Println(token_string)
	return token_string, nil
}
