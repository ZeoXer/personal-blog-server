package service

import (
	"fmt"
	"go-server/global"
	"go-server/model"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ImageService struct{}

func (i *ImageService) SaveAvatar(c *gin.Context) error {
	file, handler, err := c.Request.FormFile("uploadAvatar")
	if err != nil {
		return err
	}
	defer file.Close()

	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	serverPath := fmt.Sprintf("http://%s:%s/", global.CONFIG.Server.Host, global.CONFIG.Server.Port)

	// 檢查使用者是否有上傳過頭像，若有則刪除舊的頭像
	existingImage := model.Avatar{}
	if err := global.DB.Where("username = ?", username).First(&existingImage).Error; err == nil {
		if err := os.Remove(strings.TrimPrefix(existingImage.Path, serverPath)); err != nil {
			return err
		}
		if err := global.DB.Delete(&existingImage).Error; err != nil {
			return err
		}
	}

	// 定義圖片儲存路徑
	// 若該路徑不存在，則建立一個新的資料夾
	imgSavePath := fmt.Sprintf("uploadImgs/avatar/%s", username)
	imgPath := imgSavePath + "/" + handler.Filename
	if _, err := os.Stat(imgSavePath); os.IsNotExist(err) {
		err := os.MkdirAll(imgSavePath, 0755)
		if err != nil {
			return err
		}
	}

	// 在路徑中建立一個新的檔案並將圖片寫入
	uploadFile, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	if _, err := io.Copy(uploadFile, file); err != nil {
		return err
	}

	// 將檔案紀錄到資料庫
	image := model.Avatar{
		Username: username,
		Filename: handler.Filename,
		Path:     fmt.Sprintf("%s%s", serverPath, imgPath),
	}
	if err := global.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

func (i *ImageService) GetAvatar(c *gin.Context) (model.Avatar, error) {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return model.Avatar{}, fmt.Errorf("使用者名稱不存在")
	}

	avatar := model.Avatar{}
	if err := global.DB.Where("username = ?", username).First(&avatar).Error; err != nil {
		return model.Avatar{}, fmt.Errorf("頭像不存在")
	}

	return avatar, nil
}

func (i *ImageService) RemoveAvatar(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	serverPath := fmt.Sprintf("http://%s:%s/", global.CONFIG.Server.Host, global.CONFIG.Server.Port)

	existingImage := model.Avatar{}
	if err := global.DB.Where("username = ?", username).First(&existingImage).Error; err == nil {
		if err := os.Remove(strings.TrimPrefix(existingImage.Path, serverPath)); err != nil {
			return err
		}
		if err := global.DB.Delete(&existingImage).Error; err != nil {
			return err
		}
	}

	return nil
}

var ImageServiceGroup = new(ImageService)
