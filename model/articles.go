package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Author      string `json:"author" gorm:"not null;type:varchar(50);foreignKey:Username;"`
	CategoryId  uint   `json:"category" gorm:"not null;type:varchar(50);"`
	Title       string `json:"title" gorm:"not null;type:varchar(50);"`
	Content     string `json:"content" gorm:"type:text;"`
	IsPublished bool   `json:"is_published" gorm:"not null;default:false;"`
}

func (a Article) TableName() string {
	return "Articles"
}

type ArticleCategory struct {
	gorm.Model
	Owner    string    `json:"owner" gorm:"not null;type:varchar(50);foreignKey:Username;"`
	Category string    `json:"category" gorm:"not null;type:varchar(50);"`
	Articles []Article `json:"articles" gorm:"foreignKey:CategoryID;"`
}

func (a ArticleCategory) TableName() string {
	return "ArticleCategories"
}
