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
			// 创建Task任务：
			authed.POST("task", controller.CreateTask)
			// 查看一个Task详细信息：
			authed.GET("task/:id", controller.ShowTask)
			// 获取当前用户的所有Task任务列表：
			authed.GET("tasks", controller.ListTask)
			// 删除指定Task任务信息：
			authed.DELETE("task/:id", controller.DeleteTask)
			// 更新指定Task任务信息：
			authed.PUT("task/:id", controller.UpdateTask)
			// 搜索Task任务：
			authed.POST("search", controller.SearchTasks)
		}
	}
	return r
}
