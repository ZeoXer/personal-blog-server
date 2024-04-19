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

var ImageAPIGroup = new(ImageAPI)
