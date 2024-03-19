package model

type User struct {
	Username string `json:"username" gorm:"primaryKey;not null;type:varchar(50)"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(50)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (u User) TableName() string {
	return "Users"
}
