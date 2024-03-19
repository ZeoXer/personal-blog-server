package api

import "github.com/gin-gonic/gin"

type ImageAPI struct {
}

func (i *ImageAPI) UploadAvatar(c *gin.Context) {
	err := ImageService.SaveImage(c)
	if err != nil {
		Utils.CJSON(400, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "Image saved", nil, 1, c)
}
