package model

import (
	"awesomeProject/config/cache"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

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

func (Task *Task) AddView() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cache.RedisClient.Incr(ctx, cache.TaskViewKey(Task.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(ctx, cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
}
