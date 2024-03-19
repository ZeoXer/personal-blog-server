package global

import (
	"go-server/model"

	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	USER       model.User
	CONFIG     model.Config
	SECRET_KEY string
)
