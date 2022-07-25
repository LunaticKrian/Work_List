package database

import "awesomeProject/model"

// migration
// @Func:自动映射数据库（自动迁移模式），自动将struct映射到数据库中，自动创建数据库表
func migration() {
	Db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.User{}).AutoMigrate(&model.Task{})
	// 设置外键：
	Db.Model(&model.Task{}).AddForeignKey("uid", "user(id)", "CASCADE", "CASCADE")
}
