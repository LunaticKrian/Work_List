package service

import (
	"awesomeProject/config/utls/database"
	"awesomeProject/config/utls/serializer"
	"awesomeProject/model"
	"net/http"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

// UserRegister
// @Func: 用户注册逻辑
func (UserService *UserService) UserRegister() serializer.Response[string] {
	var user model.User
	var count int

	database.Db.Model(&model.User{}).Where("user_name=?", UserService.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response[string]{
			Status:  http.StatusBadRequest,
			Massage: "当前用户名已被注册，请更换用户名或者登录当前账户！",
		}
	}
	user.UserName = UserService.UserName

	// TODO:密码加密
	password, err := serializer.EncryptPassword(UserService.Password, 12)
	if err != nil {
		return serializer.Response[string]{
			Status:  http.StatusBadRequest,
			Massage: "服务器内部异常，密码加密失败！",
			Error:   err.Error(),
		}
	}
	user.Password = password

	// 创建用户，向数据中写入：
	if err := database.Db.Create(&user).Error; err != nil {
		return serializer.Response[string]{
			Status:  http.StatusInternalServerError,
			Massage: "数据库操作失败！",
		}
	}
	return serializer.Response[string]{
		Status:  http.StatusOK,
		Massage: "用户注册成功！",
		Data:    serializer.ToJson[model.User](user),
	}
}
