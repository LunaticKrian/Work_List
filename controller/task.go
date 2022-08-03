package controller

import (
	"awesomeProject/config/utls/jwt"
	"awesomeProject/service"
	"fmt"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

// CreateTask
// @Func: 创建一个Task任务
func CreateTask(c *gin.Context) {
	createService := service.CreateTaskService{}
	chaim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService); err == nil {
		res := createService.Create(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// ShowTask
// @Func: 查询一个Task任务的详细信息
func ShowTask(c *gin.Context) {
	showTaskService := service.ShowTaskService{}
	res := showTaskService.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// ListTask
// @Func：查询当前用户所有的Task任务
func ListTask(c *gin.Context) {
	// 引入Service：
	listService := service.ListTasksService{}
	// 验证Token，确认权限：
	chaim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		// TODO:日志输出错误信息
		fmt.Println(err.Error())
	}
}

// DeleteTask
// @Func: 删除指定的Task任务
func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// UpdateTask
// @Func: 更新Task任务信息
func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
		// TODO：日志输出异常信息
	}
}

// SearchTasks
// @Func: 条件搜索Task任务
func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	chaim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
		// TODO:日志信息
	}
}
