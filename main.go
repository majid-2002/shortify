package main

import (
	"github.com/gin-gonic/gin"
	"shortify/routes"
	"shortify/database"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetupRoutes(r, db)

	r.Run(":8080")
}