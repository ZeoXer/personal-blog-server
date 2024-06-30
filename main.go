package main

import (
	"go-server/db"
	"go-server/global"
	"go-server/router"

	"github.com/gin-contrib/cors"
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://blog.zeoxer.com"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	server.Use(cors.New(corsConfig))

	// Set the trusted proxies to nil
	server.SetTrustedProxies(nil)

	certFile := "./certificate.pem"
	keyFile := "./private-key.pem"
	server.RunTLS(":"+global.CONFIG.Server.Port, certFile, keyFile) // Run the server
}
