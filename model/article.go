package model

import (
	"time"
)

type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Username    string    `json:"username" gorm:"not null;type:varchar(50);"`
	CategoryID  uint      `json:"category_id" gorm:"not null;type:varchar(50);"`
	Title       string    `json:"title" gorm:"not null;type:varchar(50);"`
	Content     string    `json:"content" gorm:"type:text;"`
	IsPublished bool      `json:"is_published" gorm:"not null;default:false;"`
}

func (a Article) TableName() string {
	return "Articles"
}

type ArticleCategory struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string    `json:"username" gorm:"not null;type:varchar(50);"`
	CategoryName string    `json:"category_name" gorm:"not null;type:varchar(50);"`
	Articles     []Article `json:"articles" gorm:"foreignKey:CategoryID;"`
}

func (a ArticleCategory) TableName() string {
	return "ArticleCategories"
}

type ArticleAnalysis struct {
	ArticleAmount         uint `json:"article_amount"`
	ArticleCategoryAmount uint `json:"article_category_amount"`
}
