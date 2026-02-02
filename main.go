package main

import (
	"testpkg/ginserver/db"
	"testpkg/ginserver/entity"
	"testpkg/ginserver/middlewares"
	"testpkg/ginserver/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		DB setup
	*/
	db_gorm, err := db.Connect()

	if err != nil {
		panic("Failed to connect to the database")
	}

	sqlDB, err := db_gorm.DB()
	if err != nil {
		panic(err.Error())
	}
	defer sqlDB.Close()

	db_gorm.AutoMigrate(&entity.User{})

	/*
		Endpoint setup
	*/
	middlewares.SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.ErrorMiddleware())

	nonprotected := server.Group("")

	routers.SetupPublicRouters(nonprotected, db_gorm)

	protected := server.Group("")
	protected.Use(middlewares.Auth())
	{

		routers.SetupPrivateRouters(protected)
	}

	server.Run(":3000")
}
