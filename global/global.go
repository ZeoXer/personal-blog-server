package global

import (
	"go-server/model"

	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	CONFIG     model.Config
	SECRET_KEY string
)
