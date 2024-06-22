package api

import (
	"go-server/global"
	auth_model "go-server/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the tokenString from the header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			Utils.CJSON(401, "未授權", nil, 0, c)
			c.Abort()
			return
		}

		// get the token
		tokenString = tokenString[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.SECRET_KEY), nil
		})
		if err != nil {
			Utils.CJSON(401, err.Error(), nil, 0, c)
			c.Abort()
			return
		}

		// get the claims inside the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			Utils.CJSON(401, "未找到 claims", nil, 0, c)
			c.Abort()
			return
		}

		// check if the token is expired
		expired := time.Unix(int64(claims["expired"].(float64)), 0)
		if time.Now().After(expired) {
			Utils.CJSON(401, "授權已過期", nil, 0, c)
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			Utils.CJSON(401, "未找到 username", nil, 0, c)
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			Utils.CJSON(401, "未找到 email", nil, 0, c)
			c.Abort()
			return
		}

		if err := global.DB.Where("username = ? AND email = ?", username, email).First(&auth_model.User{}).Error; err != nil {
			Utils.CJSON(401, "不正確的使用者資訊", nil, 0, c)
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Set("email", email)

		c.Next()
	}
}

var AuthMiddlewareGroup = new(AuthMiddleware)
