package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/repository"
	"net/http"
)

func GetIndex(c *gin.Context) {
	repository.GetUserInfo()

	c.JSON(http.StatusOK, gin.H{
		"message": "数据表创建成功",
	})
}
