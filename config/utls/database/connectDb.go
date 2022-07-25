package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// Db 定义全局数据库对象：
var Db *gorm.DB

// DatabaseConnect
// @Func: 数据库连接
// Param: 数据库链接地址
func DatabaseConnect(databasePath string) {
	db, err := gorm.Open("mysql", databasePath)
	if err != nil {
		fmt.Println("数据库链接失败！")
	}

	// 打印输出数据库日志：
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	db.SingularTable(true)                       // 表名不加s
	db.DB().SetMaxIdleConns(20)                  // 设置连接池
	db.DB().SetMaxOpenConns(100)                 // 设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) // 最大连接时长

	Db = db

	migration()

	// TODO:日志输出：
	fmt.Println("数据库连接成功！！！")
}
