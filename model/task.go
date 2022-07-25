package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	User  User   `gorm:"ForeignKey:Uid"`
	Uid   uint   `gorm:"not null"`
	Title string `gorm:"index; not null"`

	// Task完成状态：1-表示完成;0-表示未完成：
	Status  int    `gorm:"default:'0'"`
	Content string `gorm:"type:longtext"`

	//  Task任务开始时间：
	StartTime int64

	// Task任务结束时间：
	EndTime int64
}
