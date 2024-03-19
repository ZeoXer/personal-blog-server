package router

import (
	"go-server/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	AuthAPI := api.AuthAPIGroup
	AuthMiddlewareGroup := api.AuthMiddlewareGroup

	// register the authorization routes
	authRouterGroup := router.Group("auth")
	authRouterGroup.POST("signup", AuthAPI.Signup)
	authRouterGroup.POST("login", AuthAPI.Login)

	userRouterGroup := router.Group("user")
	userRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	userRouterGroup.GET("me", AuthAPI.GetUserInfo)

	articleRouterPrivateGroup := router.Group("article")
	articleRouterPrivateGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
}
