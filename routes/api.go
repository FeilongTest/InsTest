package routes

import (
	"InsTest/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "FeiLong Test")
	})
	router.POST("/login", app.Login)
}
