package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_chat/repository"
	"net/http"
)

// GetUserInfo
// @Tags 获取用户详情
// @Accept json
// @Produce json
// @Success 200 {obj} welcome
// @Router /get.user_info [post]
func GetUserInfo(c *gin.Context) {
	user := repository.GetUserInfo()
	fmt.Println("获取用户信息", user)

	//userStr, _ := json.Marshal(user)
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "ok",
	})
}

// GetUserList
// @Tags 获取用户列表
// @Accept json
// @Produce json
// @Success 200 {obj} welcome
// @Router /get.user_list [post]
func GetUserList(c *gin.Context) {
	user := repository.GetUserList()
	fmt.Println("获取用户信息", user)
	c.JSON(http.StatusOK, gin.H{
		"user_list": user,
		"message":   "ok",
	})
}
