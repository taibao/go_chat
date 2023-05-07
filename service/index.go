package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go_chat/models"
	"go_chat/repository"
	"go_chat/service/request"
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
	var userParams request.AddUserParams
	if err := c.BindJSON(&userParams); err != nil {
		//返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("输出user", userParams.Name, userParams.Password)
	if userParams.Password != userParams.RePassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
	}
	user := models.UserBasic{}
	user.Name = userParams.Name
	user.Password = userParams.Password
	user.DeletedAt = time.Now().Local()
	user.HeartbeatTime = time.Now().Local()
	user.LogOutTime = time.Now().Local()
	user.LoginTime = time.Now().Local()

	res := repository.FindUserByName(user.Name)
	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    res,
			"message": "用户已注册",
		})
		return
	}

	repository.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"data":    "",
		"message": "创建成功",
	})
}

// DelUser
// @Summary 删除用户
// @Tags 用户模块
// @Accept json
// @param id query string false "1" 用户id
// @Produce json
// @Success 200 {obj} json{"code","message"}
// @Router /get.del_user [post]
func DelUser(c *gin.Context) {
	var userParams request.DelUserParams
	if err := c.BindJSON(&userParams); err != nil {
		//返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.UserBasic
	user.ID = userParams.Id
	fmt.Println("user", user)
	repository.DelUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"data":    "",
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @Accept json
// @param id formData string false "1"
// @param name formData string false "taibao"
// @param password formData string false "123123"
// @param phone formData string false "123123"
// @Produce json
// @Success 200 {obj} json{"code","message"}
// @Router /update_user [post]
func UpdateUser(c *gin.Context) {
	var user models.UserBasic
	if err := c.BindJSON(&user); err != nil {
		//返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user", user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    "0",
		"message": "更新成功",
	})
}
