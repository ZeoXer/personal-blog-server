package service

import (
	"fmt"
	"go-server/global"
	"go-server/model"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageService struct{}

func (i *ImageService) SaveAvatar(c *gin.Context) error {
	file, handler, err := c.Request.FormFile("uploadAvatar")
	if err != nil {
		return err
	}
	defer file.Close()

	// Save the image
	imgPath := fmt.Sprintf("d/uploadImgs/avatar/%s/%s", global.USER.Username, handler.Filename)
	uploadFile, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	// Create a new Image record
	image := model.Avatar{
		Username: global.USER.Username,
		Filename: handler.Filename,
		Path:     imgPath,
	}

	// Save the Image record to the database
	if err := global.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

var ImageServiceGroup = new(ImageService)
