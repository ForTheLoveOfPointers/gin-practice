package main

import (
	"database/sql"
	"log"
	"testpkg/ginserver/db"
	"testpkg/ginserver/middlewares"
	"testpkg/ginserver/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT);")

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db_gorm, err := db.Connect()

	if err != nil {
		panic("Failed to connect to the database")
	}

	middlewares.SetupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.ErrorMiddleware())

	nonprotected := server.Group("/basic")

	routers.SetupPublicRouters(nonprotected)

	protected := server.Group("/my-account")
	protected.Use(middlewares.Auth())
	{

		routers.SetupPrivateRouters(protected)
	}

	server.Run(":3000")
}
