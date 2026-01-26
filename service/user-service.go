package service

import (
	"fmt"
	"slices"
	"testpkg/ginserver/entity"
	servererrors "testpkg/ginserver/server_errors"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Register(entity.User) entity.User
	Login(entity.User) (string, error)
}

type userService struct {
	users []entity.User
}

func NewUser() *userService {
	return &userService{}
}

func (e *userService) Register(usr entity.User) entity.User {
	e.users = append(e.users, usr)
	return usr
}

func (e *userService) Login(usr entity.User) (string, error) {
	if !slices.Contains(e.users, usr) {
		err := servererrors.RequestError{Code: 404, Message: "User not in database"}
		return "", &err
	}
	fmt.Println(usr)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":  usr.Name,
			"email": usr.Email,
		},
	)

	token_string, err := token.SignedString([]byte("MYSECRET"))
	if err != nil {
		return "", &servererrors.RequestError{Code: 500, Message: "Failed at jwt signing"}
	}

	fmt.Println(token_string)
	return token_string, nil
}
