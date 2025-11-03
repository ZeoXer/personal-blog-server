package service

import (
	"context"
	"fmt"
	"go-server/global"
	"go-server/model"
	"go-server/r2client"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	serverPath := fmt.Sprintf("https://%s/", global.CONFIG.Server.Host)

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
	authorName := c.Param("authorName")
	username, _, err := Utils.GetUserInfo(c)
	if err != nil && authorName == "" {
		return model.Avatar{}, fmt.Errorf("使用者名稱不存在")
	}

	avatar := model.Avatar{}
	searchName := username
	if authorName != "" {
		searchName = authorName
	}

	if err := global.DB.Where("username = ?", searchName).First(&avatar).Error; err != nil {
		return model.Avatar{}, fmt.Errorf("頭像不存在")
	}

	return avatar, nil
}

func (i *ImageService) RemoveAvatar(c *gin.Context) error {
	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return err
	}

	serverPath := fmt.Sprintf("https://%s/", global.CONFIG.Server.Host)

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

func (i *ImageService) SaveImage(c *gin.Context) (model.Image, error) {
	image := model.Image{}

	file, handler, err := c.Request.FormFile("uploadImage")
	if err != nil {
		return image, err
	}
	defer file.Close()

	username, _, err := Utils.GetUserInfo(c)
	if err != nil {
		return image, err
	}

	serverPath := fmt.Sprintf("https://%s/", global.CONFIG.Server.Host)
	imgSavePath := fmt.Sprintf("uploadImgs/image/%s", username)
	imgPath := imgSavePath + "/" + uuid.New().String() + ".png"
	if _, err := os.Stat(imgSavePath); os.IsNotExist(err) {
		err := os.MkdirAll(imgSavePath, 0755)
		if err != nil {
			return image, err
		}
	}

	uploadFile, err := os.OpenFile(imgPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return image, err
	}
	defer uploadFile.Close()

	if _, err := io.Copy(uploadFile, file); err != nil {
		return image, err
	}

	image = model.Image{
		Filename: handler.Filename,
		Path:     fmt.Sprintf("%s%s", serverPath, imgPath),
	}
	if err := global.DB.Create(&image).Error; err != nil {
		return image, err
	}

	return image, nil
}

func (i *ImageService) ListObjectsR2(c *gin.Context) ([]gin.H, error) {
	limit := int32(1000)
	bucket := global.CONFIG.R2Storage.BucketName

	ctx := c.Request.Context()
	client, err := r2client.NewR2Client(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := client.S3.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  &bucket,
		MaxKeys: &limit,
	})

	if err != nil {
		return nil, err
	}

	items := make([]gin.H, 0, len(res.Contents))
	for _, obj := range res.Contents {
		items = append(items, gin.H{
			"key":           *obj.Key,
			"size":          obj.Size,
			"last_modified": obj.LastModified,
			"etag":          obj.ETag,
			"storage_class": obj.StorageClass,
		})
	}

	return items, nil
}

func getContentType(header *multipart.FileHeader, name string) string {
	if v := header.Header.Get("Content-Type"); v != "" {
		return v
	}
	switch filepath.Ext(name) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}

func (i *ImageService) PutObjectR2(c *gin.Context) (string, error) {
	if err := c.Request.ParseMultipartForm(64 << 20); err != nil {
		return "", err
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	if fileHeader.Size > (200 << 20) {
		return "", fmt.Errorf("file size exceeds 200MB limit")
	}

	f, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	filenameSafe := regexp.MustCompile(`[^a-zA-Z0-9._-]+`)
	safe := filenameSafe.ReplaceAllString(fileHeader.Filename, "_")
	contentType := getContentType(fileHeader, safe)
	contentDisposition := fmt.Sprintf(`inline; filename="%s"`, filepath.Base(safe))

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Minute)
	defer cancel()

	client, err := r2client.NewR2Client(ctx)
	bucket := global.CONFIG.R2Storage.BucketName
	if err != nil {
		return "", err
	}
	uploader := manager.NewUploader(client.S3)

	out, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:             &bucket,
		Key:                &safe,
		Body:               f,
		ContentType:        &contentType,
		ContentDisposition: &contentDisposition,
		Metadata: map[string]string{
			"uploadedBy": "gin-backend",
			"size":       strconv.FormatInt(fileHeader.Size, 10),
		},
	})
	if err != nil {
		return "", err
	}

	return out.Location, nil
}

var ImageServiceGroup = new(ImageService)
