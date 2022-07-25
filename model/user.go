package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`

	// 使用加密后的密文存储到数据库：
	Password string
}
