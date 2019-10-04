package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type UserPicture struct {
	gorm.Model
	User           User
	Picture        Picture
	UserID         int64          `gorm:"unique_index:user_picture_u1;not null"`
	PictureID      int64          `gorm:"unique_index:user_picture_u1;not null"`
	Name           string         `gorm:"type:varchar(255);index;not null"`
	Score          int            `gorm:"index;default:'0'"`
	IsPublic       bool           `gorm:"default:'0'"`
	PublicToken    sql.NullString `gorm:"type:varchar(64);index"`
	UserPictureTag []Tag          `gorm:"many2many:user_picture_tag"`
}
