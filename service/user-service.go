package service

import (
	"fmt"
	"os"
	"testpkg/ginserver/entity"
	servererrors "testpkg/ginserver/server_errors"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserService interface {
	Register(entity.User) (string, error)
	Login(entity.User) (string, error)
}

type userService struct {
	db *gorm.DB
}

func signToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func NewUser(db *gorm.DB) *userService {
	return &userService{db}
}

func (e *userService) Register(usr entity.User) (string, error) {
	result := e.db.Create(&usr)
	if result.Error != nil {
		err := servererrors.RequestError{Code: 500, Message: "Error registering user"}
		return "", &err
	}
	token_string, err := signToken(jwt.MapClaims{
		"email": usr.Email,
		"name":  usr.Name,
	})

	if err != nil {
		return "", &servererrors.RequestError{Code: 500, Message: "Failed at jwt signing"}
	}

	return token_string, nil
}

func (e *userService) Login(usr entity.User) (string, error) {
	var userEnt entity.User
	result := e.db.Where("email = ?", usr.Email).First(&userEnt)

	if result.Error != nil {
		err := servererrors.RequestError{Code: 500, Message: "User unable to login"}
		return "", &err
	}

	token_string, err := signToken(jwt.MapClaims{
		"name":  userEnt.Name,
		"email": userEnt.Email,
	})

	if err != nil {
		return "", &servererrors.RequestError{Code: 500, Message: "Failed at jwt signing"}
	}

	fmt.Println(token_string)
	return token_string, nil
}
