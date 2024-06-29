package router

import (
	"go-server/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// register api group
	AuthAPI := api.AuthAPIGroup
	ImageAPI := api.ImageAPIGroup
	ArticleAPI := api.ArticleAPIGroup
	AuthMiddlewareGroup := api.AuthMiddlewareGroup

	// /auth/...
	authRouterGroup := router.Group("auth")
	authRouterGroup.POST("signup", AuthAPI.Signup)
	authRouterGroup.POST("login", AuthAPI.Login)

	// /user/...
	userRouterGroup := router.Group("user")
	userRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	userRouterGroup.GET("me", AuthAPI.GetUserInfo)

	// /article/...
	articleRouterPrivateGroup := router.Group("article")
	articleRouterPrivateGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	articleRouterPrivateGroup.POST("addArticleCategory", ArticleAPI.AddArticleCategory)
	articleRouterPrivateGroup.GET("getAllArticleCategory", ArticleAPI.GetAllArticleCategory)
	articleRouterPrivateGroup.PUT("updateArticleCategory/:categoryId", ArticleAPI.UpdateArticleCategory)
	articleRouterPrivateGroup.POST("addArticle", ArticleAPI.AddArticle)
	articleRouterPrivateGroup.GET("getArticle/:articleId", ArticleAPI.GetArticle)
	articleRouterPrivateGroup.PUT("updateArticle/:articleId", ArticleAPI.UpdateArticle)
	articleRouterPrivateGroup.GET("getArticlesByCategory/:categoryId", ArticleAPI.GetArticlesByCategory)

	// /image/...
	imgRouterGroup := router.Group("image")
	// 定義靜態資源路徑
	router.Static("/uploadImgs/avatar", "./uploadImgs/avatar")
	imgRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	imgRouterGroup.POST("uploadAvatar", ImageAPI.UploadAvatar)
	imgRouterGroup.GET("getAvatar", ImageAPI.GetAvatar)
	imgRouterGroup.DELETE("removeAvatar", ImageAPI.RemoveAvatar)
}
