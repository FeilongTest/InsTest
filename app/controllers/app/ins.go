package app

import (
	"InsTest/app/common/request"
	"InsTest/app/common/response"
	"InsTest/app/services/ins"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	//账号----密码
	var form request.LoginData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	temp, err := ins.Login(form)
	if err == nil {
		response.Success(c, temp)
	} else {
		response.BusinessFail(c, err.Error())
	}
}

//SendMessage 发送私信
func SendMessage(c *gin.Context) {
	var form request.MessageData
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	temp, err := ins.SendMessage(form)
	if err == nil {
		response.Success(c, temp)
	} else {
		response.BusinessFail(c, err.Error())
	}
}
