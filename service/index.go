package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_chat/models"
	"go_chat/repository"
	"net/http"
	"time"
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

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Accept json
// @param name query string false "用户名"
// @param password query string false "密码"
// @param re_password query string false "确认密码"
// @Produce json
// @Success 200 {obj} json{"code","message"}
// @Router /get.add_user [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	rePassword := c.Query("re_password")

	if password != rePassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}
	user.Password = password
	user.DeletedAt = time.Now().Local()
	user.HeartbeatTime = time.Now().Local()
	user.LogOutTime = time.Now().Local()
	user.LoginTime = time.Now().Local()

	repository.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"data":    "",
		"message": "创建成功",
	})
}
