package router

import (
	"go-server/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// register api group
	AuthAPI := api.AuthAPIGroup
	ImageAPI := api.ImageAPIGroup
	AuthMiddlewareGroup := api.AuthMiddlewareGroup

	// register routes by group
	authRouterGroup := router.Group("auth")
	authRouterGroup.POST("signup", AuthAPI.Signup)
	authRouterGroup.POST("login", AuthAPI.Login)

	userRouterGroup := router.Group("user")
	userRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	userRouterGroup.GET("me", AuthAPI.GetUserInfo)

	articleRouterPrivateGroup := router.Group("article")
	articleRouterPrivateGroup.Use(AuthMiddlewareGroup.AuthMiddleware())

	imgRouterGroup := router.Group("image")
	imgRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	imgRouterGroup.POST("uploadAvatar", ImageAPI.UploadAvatar)
}
