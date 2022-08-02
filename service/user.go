package service

import (
	"awesomeProject/config/utls/database"
	"awesomeProject/config/utls/jwt"
	"awesomeProject/config/utls/serializer"
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
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
			Message: "当前用户名已被注册，请更换用户名或者登录当前账户！",
		}
	}
	user.UserName = UserService.UserName

	// TODO:密码加密
	password, err := serializer.EncryptPassword(UserService.Password, 12)
	if err != nil {
		return serializer.Response[string]{
			Status:  http.StatusBadRequest,
			Message: "服务器内部异常，密码加密失败！",
			Error:   err.Error(),
		}
	}
	user.Password = password

	// 创建用户，向数据中写入：
	if err := database.Db.Create(&user).Error; err != nil {
		return serializer.Response[string]{
			Status:  http.StatusInternalServerError,
			Message: "数据库操作失败！",
		}
	}
	return serializer.Response[string]{
		Status:  http.StatusOK,
		Message: "用户注册成功！",
		Data:    serializer.ToJson[model.User](user),
	}
}

// UserLogin
// @Func: 用户登录逻辑
func (UserService *UserService) UserLogin() serializer.Response[interface{}] {
	var user model.User

	// 从数据库中查询：
	if err := database.Db.Where("user_name=?", UserService.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response[interface{}]{
				Status:  http.StatusBadRequest,
				Message: "当前登录用户不存在，请注册后登录！",
			}
		}
		// 其他错误情况：
		return serializer.Response[interface{}]{
			Status:  http.StatusInternalServerError,
			Message: "服务器发生了未知错误！",
		}
	}

	// 校验密码：
	if serializer.ComparePassword(user.Password, UserService.Password) == false {
		return serializer.Response[interface{}]{
			Status:  http.StatusBadRequest,
			Message: "密码输入错误！",
		}
	}

	// 发送Token，身份验证：
	token, err := jwt.GenerateToken(user.ID, user.UserName, user.Password)
	if err != nil {
		return serializer.Response[interface{}]{
			Status:  http.StatusInternalServerError,
			Message: "token签发错误！！！",
		}
	}

	// 正常返回：
	return serializer.Response[interface{}]{
		Status:  http.StatusOK,
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Message: "登录成功！",
	}
}
