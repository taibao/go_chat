package router

import (
	"github.com/gin-gonic/gin"
	"go_chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	return r
}
