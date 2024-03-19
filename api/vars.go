package api

import (
	service "go-server/service"
	"go-server/utils"
)

var (
	Utils          = utils.Utils
	AuthService    = service.AuthServiceGroup
	ArticleService = service.ArticleServiceGroup
	ImageService   = service.ImageServiceGroup
)
