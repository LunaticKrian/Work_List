package controller

import (
	"awesomeProject/config/utls/jwt"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

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
