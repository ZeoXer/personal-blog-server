package utils

import "github.com/gin-gonic/gin"

type Response struct{}

func (r *Response) CJSON(code int, message string, result any, status int, c *gin.Context) {
	c.JSON(code, gin.H{
		"data":    result,
		"message": message,
		"status":  status,
	})
}

var Utils = new(Response)
