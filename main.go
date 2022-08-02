package main

import (
	config "awesomeProject/config/local"
	"awesomeProject/routes"
)

func main() {
	// 项目服务启动初始化：
	config.Init()

	// 创建根路由：
	r := routes.NewRouter()
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
