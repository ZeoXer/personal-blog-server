package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UtilFunc struct{}

func (u *UtilFunc) CJSON(code int, message string, result any, status int, c *gin.Context) {
	c.JSON(code, gin.H{
		"data":    result,
		"message": message,
		"status":  status,
	})
}

// 回傳格式(username, email, error)
func (u *UtilFunc) GetUserInfo(c *gin.Context) (string, string, error) {
	username, exists := c.Get("username")
	if !exists {
		return "", "", fmt.Errorf("找不到使用者")
	}

	email, exists := c.Get("email")
	if !exists {
		return "", "", fmt.Errorf("找不到信箱")
	}

	return username.(string), email.(string), nil
}

var Utils = new(UtilFunc)
