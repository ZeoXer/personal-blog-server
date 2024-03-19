package model

type Avatar struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null;type:varchar(50);foreignKey:Username"`
	Filename string `json:"filename" gorm:"not null;type:varchar(255)"`
	Path     string `json:"path" gorm:"not null;type:varchar(255)"`
}
