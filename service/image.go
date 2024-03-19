package service

import (
	"go-server/global"
	"go-server/model"

	"github.com/gin-gonic/gin"
)

type ImageService struct{}

func (i *ImageService) SaveImage(c *gin.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	// Save the image to db
	imagePath := "images/" + file.Filename
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		return err
	}

	// Create a new Image record
	image := model.Avatar{
		Username: global.USER.Username,
		Filename: file.Filename,
		Path:     imagePath,
	}

	// Save the Image record to the database
	if err := global.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

var ImageServiceGroup = new(ImageService)
