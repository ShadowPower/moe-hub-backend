package model

type Tag struct {
	Name string `gorm:"type:varchar(64);primary_key"`
}
