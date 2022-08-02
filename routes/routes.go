package routes

import (
	"awesomeProject/config/utls/midlleware"
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
		v1.POST("user/login", controller.UserLogin)

		authed := v1.Group("/")
		authed.Use(midlleware.JWT())
		{
			// 任务操作：
			authed.POST("task", controller.CreateTask)
		}
	}
	return r
}
