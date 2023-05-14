package service

import (
	"github.com/gin-gonic/gin"
	"go_chat/models"
	"go_chat/utils"
)

func CreateTable(c *gin.Context) {
	{
		db := utils.DB
		//user := models.UserBasic{}
		//db.Where("id", 1).Find(&user)
		//
		//return user
		// 迁移 schema
		db.AutoMigrate(&models.Contact{})
		db.AutoMigrate(&models.GroupBasic{})
		db.AutoMigrate(&models.Message{})

		// Create
		//user := &models.UserBasic{}
		//user.Name = "vitas"
		//user.LogOutTime = time.Now().Local()
		//user.LoginTime = time.Now().Local()
		//user.HeartbeatTime = time.Now().Local()
		//db.Create(user)
		//db.Where("Name = ?", "vitas").First(user)
		//fmt.Println("输出第一条记录", user)

		// Read
		//db.First(user, 1)                 // 根据整型主键查找
		//db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

		// Update - 将 product 的 price 更新为 200
		//db.Model(user).Update("PassWord", "1234")
		//// Update - 更新多个字段
		//db.Model(user).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
		//db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
		//
		//// Delete - 删除 product
		//db.Delete(user, 1)
	}

}
