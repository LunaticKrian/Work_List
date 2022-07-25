package controller

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
