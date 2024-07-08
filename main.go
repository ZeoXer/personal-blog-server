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

	certFile := "./certificate.pem"
	keyFile := "./private-key.pem"
	server.RunTLS(":"+global.CONFIG.Server.Port, certFile, keyFile) // Dep

	// server.Run(":" + global.CONFIG.Server.Port) // Dev
}
