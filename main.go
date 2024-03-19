package main

import (
	"go-server/db"
	"go-server/global"
	"go-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	var err error

	// Read the config file
	InitViper()

	// Connect to the database and automigrate
	global.DB, err = db.Connect()
	if err != nil {
		panic("Error connecting db: " + err.Error())
	}
	db.AutoMigrate()

	// Initialize the routes
	router.InitializeRoutes(server)

	// Set the trusted proxies to nil
	server.SetTrustedProxies(nil)

	server.Run(":" + global.CONFIG.Server.Port) // Run the server
}
