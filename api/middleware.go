package api

import (
	"go-server/global"
	"go-server/model"
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
			Utils.CJSON(401, "No claims found", nil, 0, c)
			c.Abort()
			return
		}

		// check if the token is expired
		expired := time.Unix(int64(claims["expired"].(float64)), 0)
		if time.Now().After(expired) {
			Utils.CJSON(401, "Token expired", nil, 0, c)
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			Utils.CJSON(401, "Username not found", nil, 0, c)
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			Utils.CJSON(401, "Email not found", nil, 0, c)
			c.Abort()
			return
		}

		global.USER = model.User{
			Username: username,
			Email:    email,
		}

		c.Next()
	}
}

var AuthMiddlewareGroup = new(AuthMiddleware)
