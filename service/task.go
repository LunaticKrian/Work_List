package service

import (
	"awesomeProject/config/utls/database"
	"awesomeProject/config/utls/serializer"
	"awesomeProject/model"
	"net/http"
	"time"
)

// CreateTaskService
// @Func: 接收前端提交的JSON（参数），创建Task任务的信息
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

// ShowTaskService
// @Func：Task详细信息
type ShowTaskService struct {
}

// ListTasksService
// @Func: 接收前端请求参数，Task任务信息列表
type ListTasksService struct {
	// 分页信息：
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

// DeleteTaskService
// @Func：接收前端参数，指定Task
type DeleteTaskService struct {
}

// UpdateTaskService
// @Func: 更新Task任务信息
type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

// SearchTaskService
// @Func：接收前端提交的搜索信息
type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

// Create
// @Func：创建Task任务
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

// Show
// @Func：查询Task任务详细信息
func (service *ShowTaskService) Show(id string) serializer.Response[interface{}] {
	var task model.Task
	code := http.StatusOK
	err := database.Db.First(&task, id).Error
	if err != nil {
		code = http.StatusInternalServerError
		return serializer.Response[interface{}]{
			Status:  code,
			Message: "服务器发生内部错误！",
			Error:   err.Error(),
		}
	}
	task.AddView() //增加点击数
	return serializer.Response[interface{}]{
		Status:  code,
		Data:    task,
		Message: "获取Task信息成功！",
	}
}

// List
// @Func：查询当前用户的所有Task任务
func (service *ListTasksService) List(id uint) serializer.Response[interface{}] {
	var tasks []model.Task
	var total int64
	// 设置默认的分页参数：
	if service.Limit == 0 {
		service.Limit = 15
	}
	database.Db.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serializer.Response[interface{}]{
		Status:  http.StatusOK,
		Message: "查询任务列表成功！",
		Data:    tasks,
	}
}

// Delete
// @Func：删除指定的Task任务
func (service *DeleteTaskService) Delete(id string) serializer.Response[interface{}] {
	var task model.Task
	code := http.StatusOK
	// 查询数据中是否含由当前删除的Task任务：
	err := database.Db.First(&task, id).Error
	if err != nil {
		// 数据库查询异常：
		// TODO：日志输出异常
		return serializer.Response[interface{}]{
			Status:  http.StatusInternalServerError,
			Message: "服务器内部错误！",
			Error:   err.Error(),
		}
	}
	err = database.Db.Delete(&task).Error
	if err != nil {
		// 数据库删除异常：
		// TODO：日志输出异常信息
		return serializer.Response[interface{}]{
			Status:  http.StatusInternalServerError,
			Message: "服务器内部错误！",
			Error:   err.Error(),
		}
	}
	return serializer.Response[interface{}]{
		Status:  code,
		Message: "任务删除成功！",
	}
}

func (service *UpdateTaskService) Update(id string) serializer.Response[interface{}] {
	var task model.Task
	database.Db.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := http.StatusOK
	err := database.Db.Save(&task).Error
	if err != nil {
		return serializer.Response[interface{}]{
			Status:  code,
			Message: "服务器内部错误！",
			Error:   err.Error(),
		}
	}
	return serializer.Response[interface{}]{
		Status:  code,
		Message: "Task任务信息更新成功！",
		Data:    task,
	}
}

// Search
// @Func：按照指定信息搜索Task任务
func (service *SearchTaskService) Search(uId uint) serializer.Response[interface{}] {
	var tasks []model.Task
	code := http.StatusOK
	database.Db.Where("uid=?", uId).Preload("User").First(&tasks)
	err := database.Db.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Find(&tasks).Error
	if err != nil {
		// TODO: 日志信息输出
		return serializer.Response[interface{}]{
			Status:  http.StatusInternalServerError,
			Message: "服务器内部错误！",
			Error:   err.Error(),
		}
	}
	return serializer.Response[interface{}]{
		Status:  code,
		Message: "查询成功！",
		Data:    tasks,
	}
}
