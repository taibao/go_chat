package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_chat/docs"
	"go_chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/get.user_info", service.GetUserInfo)
	r.GET("/get.user_list", service.GetUserList)
	r.POST("/add_user", service.CreateUser)
	r.POST("/del_user", service.DelUser)
	r.POST("/update_user", service.UpdateUser)
	return r
}
