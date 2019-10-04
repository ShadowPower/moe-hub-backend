package model

import "github.com/jinzhu/gorm"

type Picture struct {
	gorm.Model
	DataHash       string        `gorm:"type:varchar(64);unique_index;not null"`
	PerceptionHash string        `gorm:"type:varchar(64);unique_index;not null"`
	DataLength     int64         `gorm:"not null"`
	UserPicture    []UserPicture `gorm:"foreignkey:PictureID"`
	Width          int           `gorm:"not null"`
	Height         int           `gorm:"not null"`
	Ratio          float64       `gorm:"not null"`
}
