package api

import (
	"github.com/gin-gonic/gin"
)

type ImageAPI struct {
}

func (i *ImageAPI) UploadAvatar(c *gin.Context) {
	err := ImageService.SaveAvatar(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "圖片上傳成功", nil, 1, c)
}

func (i *ImageAPI) GetAvatar(c *gin.Context) {
	avatar, err := ImageService.GetAvatar(c)

	if err != nil {
		Utils.CJSON(200, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "取得頭像成功", avatar, 1, c)
}

func (i *ImageAPI) RemoveAvatar(c *gin.Context) {
	err := ImageService.RemoveAvatar(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "刪除頭像成功", nil, 1, c)
}

func (i *ImageAPI) UploadImage(c *gin.Context) {
	image, err := ImageService.SaveImage(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), image, 0, c)
		return
	}

	Utils.CJSON(200, "圖片上傳成功", image, 1, c)
}

func (i *ImageAPI) ListObjectsR2(c *gin.Context) {
	objects, err := ImageService.ListObjectsR2(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "取得物件列表成功", objects, 1, c)
}

func (i *ImageAPI) PutObjectR2(c *gin.Context) {
	url, err := ImageService.PutObjectR2(c)

	if err != nil {
		Utils.CJSON(500, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "上傳物件成功", url, 1, c)
}

var ImageAPIGroup = new(ImageAPI)
