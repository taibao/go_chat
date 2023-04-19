package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/get.user_info", service.GetUserInfo)
	r.GET("/get.user_list", service.GetUserList)
	return r
}
