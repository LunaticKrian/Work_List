package config

import (
	"awesomeProject/config/utls/database"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

// config.ini 配置文件变量：
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// Init
// @Func:初始化函数，引入包时直接调用
func Init() {
	// 使用ini包下的Load函数加载配置文件：(注意路径开始是工程根目录)
	conf, err := ini.Load("./config/local/config.ini")

	// TODO:这里需要修改，引入日志框架：
	if err != nil {
		fmt.Println("配置文件读取失败！")
	}
	LoadServer(conf)
	LoadMySQL(conf)

	// 对数据库链接地址进行拼接：
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")

	database.DatabaseConnect(path)
}

// LoadServer
// @Func：解析server配置信息
// @Param: 配置文件指针
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
}

// LoadMySQL
// @Func：解析MySQL配置信息
// @Param: 配置文件指针
func LoadMySQL(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
