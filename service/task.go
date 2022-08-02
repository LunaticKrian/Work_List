package service

import (
	"awesomeProject/config/utls/database"
	"awesomeProject/config/utls/serializer"
	"awesomeProject/model"
	"net/http"
	"time"
)

type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

func (service *CreateTaskService) Create(id uint) serializer.Response[interface{}] {
	var user model.User
	database.Db.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	code := http.StatusOK
	err := database.Db.Create(&task).Error
	if err != nil {
		code = http.StatusInternalServerError
		return serializer.Response[interface{}]{
			Status:  code,
			Message: "创建任务失败！服务器内部错误！",
		}
	}
	return serializer.Response[interface{}]{
		Status:  code,
		Message: "创建任务成功！",
	}
}
