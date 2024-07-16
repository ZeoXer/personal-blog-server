package router

import (
	"go-server/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// setting cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://blog.zeoxer.com", "http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))

	router.OPTIONS("/*cors", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://blog.zeoxer.com")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.AbortWithStatus(204)
	})

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
	articleRouterPrivateGroup.DELETE("deleteArticleCategory/:categoryId", ArticleAPI.DeleteArticleCategory)
	articleRouterPrivateGroup.POST("addArticle", ArticleAPI.AddArticle)
	articleRouterPrivateGroup.GET("getArticle/:articleId", ArticleAPI.GetArticle)
	articleRouterPrivateGroup.PUT("updateArticle/:articleId", ArticleAPI.UpdateArticle)
	articleRouterPrivateGroup.DELETE("deleteArticle/:articleId", ArticleAPI.DeleteArticle)
	articleRouterPrivateGroup.GET("getArticlesByCategory/:categoryId", ArticleAPI.GetArticlesByCategory)
	articleRouterPrivateGroup.GET("getArticleAnalysis", ArticleAPI.GetArticleAnalysis)

	// /article/public/...
	articleRouterPublicGroup := router.Group("article/public")
	articleRouterPublicGroup.GET("getAllArticleCategory/:authorName", ArticleAPI.GetAllArticleCategory)
	articleRouterPublicGroup.GET("getArticlesByCategory/:authorName/:categoryId", ArticleAPI.GetArticlesByCategory)
	articleRouterPublicGroup.GET("getArticle/:authorName/:articleId", ArticleAPI.GetArticle)

	// /image/...
	imgRouterGroup := router.Group("image")
	// 定義靜態資源路徑
	router.Static("/uploadImgs/avatar", "./uploadImgs/avatar")
	imgRouterGroup.Use(AuthMiddlewareGroup.AuthMiddleware())
	imgRouterGroup.POST("uploadAvatar", ImageAPI.UploadAvatar)
	imgRouterGroup.GET("getAvatar", ImageAPI.GetAvatar)
	imgRouterGroup.DELETE("removeAvatar", ImageAPI.RemoveAvatar)
	imgRouterGroup.POST("uploadImage", ImageAPI.UploadImage)

	// /image/public/...
	imgPublicRouterGroup := router.Group("image/public")
	imgPublicRouterGroup.GET("getAvatar/:authorName", ImageAPI.GetAvatar)
}
