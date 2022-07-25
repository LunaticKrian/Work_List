package routes

import (
	"awesomeProject/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mySession", store))

	v1 := r.Group("api/v1")
	{
		// 用户操作：
		v1.POST("user/register", controller.UserRegister)
	}
	return r
}
