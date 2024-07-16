package model

type Avatar struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null;type:varchar(50);"`
	Filename string `json:"filename" gorm:"not null;type:varchar(255)"`
	Path     string `json:"path" gorm:"not null;type:varchar(255)"`
}

func (a Avatar) TableName() string {
	return "Avatars"
}

type Image struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Filename string `json:"filename" gorm:"not null;type:varchar(255)"`
	Path     string `json:"path" gorm:"not null;type:varchar(255)"`
}

func (i Image) TableName() string {
	return "Images"
}
