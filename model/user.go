package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(16);unique_index;not null"`
	Email       string `gorm:"type:varchar(254);unique_index;not null"`
	Password    string `gorm:"not null"`
	LastLoginAt *time.Time
	UserPicture []UserPicture `gorm:"foreignkey:UserID"`
}
