package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"go-server/global"
	auth_model "go-server/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthService struct{}

func getHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func getJWTToken(username, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"expired":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Generate a random string as the signature key
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes: ", err)
		return ""
	}
	hash := sha256.Sum256(randomBytes)
	tokenSignature := base64.URLEncoding.EncodeToString(hash[:])

	// Store the signature key in global constant
	global.SECRET_KEY = tokenSignature

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(tokenSignature))
	if err != nil {
		fmt.Println("Error signing token: ", err)
		return ""
	}

	return tokenString
}

func (a *AuthService) Signup(username, email, password string) (string, error) {
	var user = auth_model.User{
		Username: username,
		Email:    email,
		Password: getHashPassword(password),
	}

	// Check if the username already exists in the database
	err := global.DB.Where("username = ?", username).First(&auth_model.User{}).Error
	if err != gorm.ErrRecordNotFound {
		// Username exists in the database, return error
		return "username", errors.New("使用者名稱已被註冊")
	}

	// Check if the email already exists in the database
	err = global.DB.Where("email = ?", email).First(&auth_model.User{}).Error
	if err != gorm.ErrRecordNotFound {
		// Email exists in the database, return error
		return "email", errors.New("信箱已被註冊")
	}

	// Create the user in the database
	err = global.DB.Create(user).Error
	if err != nil {
		return "create user error", err
	}

	return "success", nil
}

func (a *AuthService) Login(email, password string) (*auth_model.LoginResponse, error) {
	var user *auth_model.User
	err := global.DB.Where("email = ? AND password = ?", email, getHashPassword(password)).First(&user).Error
	if err != nil {
		return nil, errors.New("帳號或密碼錯誤")
	}

	return &auth_model.LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    getJWTToken(user.Username, email),
	}, nil
}

var AuthServiceGroup = new(AuthService)
