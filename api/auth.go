package api

import (
	auth_model "go-server/model"

	"github.com/gin-gonic/gin"
)

type AuthAPI struct {
}

func (a *AuthAPI) Signup(c *gin.Context) {
	var user *auth_model.User
	var errorTerm string

	err := c.ShouldBindJSON(&user)
	if err != nil {
		Utils.CJSON(400, "參數缺失", nil, 0, c)
		return
	}

	errorTerm, err = AuthService.Signup(user.Username, user.Email, user.Password)
	if err != nil {
		Utils.CJSON(200, err.Error(), errorTerm, 0, c)
		return
	}

	Utils.CJSON(200, "註冊成功", nil, 1, c)
}

func (a *AuthAPI) Login(c *gin.Context) {
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		Utils.CJSON(400, "參數缺失", nil, 0, c)
		return
	}

	response, err := AuthService.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		Utils.CJSON(401, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "登入成功", response, 1, c)
}

func (a *AuthAPI) GetUserInfo(c *gin.Context) {
	username, email, err := Utils.GetUserInfo(c)
	if err != nil {
		Utils.CJSON(401, err.Error(), nil, 0, c)
		return
	}

	Utils.CJSON(200, "回傳使用者", gin.H{
		"username": username,
		"email":    email,
	}, 1, c)
}

var AuthAPIGroup = new(AuthAPI)
