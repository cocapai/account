package model

import "github.com/jinzhu/gorm"

type User struct {
	// gorm.Model嵌入数据库ID和增删改时间
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}