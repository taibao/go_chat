package repository

import (
	"fmt"
	"go_chat/models"
	"go_chat/utils"
	"gorm.io/gorm"
	"time"
)

func GetUserInfo(name, password string) *models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("name", name).Where("password", password).First(&user)
	return &user
}

func GetUserList() []models.UserBasic {
	user := make([]models.UserBasic, 2)
	utils.DB.Find(&user)
	return user
}

func FindUserByName(name string) *models.UserBasic {
	user := models.UserBasic{}
	utils.DB.Where("name", name).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)

	return &user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := models.UserBasic{}
	return utils.DB.Where("phone", phone).First(&user)
}

func CreateUser(user *models.UserBasic) bool {
	utils.DB.Create(user)
	return true
}

func FindUserByEmail(email string) *gorm.DB {
	user := models.UserBasic{}
	return utils.DB.Where("email", email).First(&user)
}

func DelUser(user *models.UserBasic) bool {
	utils.DB.Delete(user)
	return true
}

func UpdateUser(user *models.UserBasic) bool {
	utils.DB.Model(&user).Updates(models.UserBasic{Name: user.Name, Phone: user.Phone, Password: user.Password, Email: user.Email})
	return true
}
