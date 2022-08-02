package controller

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister
// @Func:用户注册方法
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	// 绑定服务对象：
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.UserRegister()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// UserLogin
// @Func: 用户登录方法
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	// 绑定服务对象：
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.UserLogin()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
