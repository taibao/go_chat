package repository

import (
	"go_chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetUserInfo() {
	db, err := gorm.Open(mysql.Open("root:123123@tcp(localhost:3306)/go_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "vitas"
	user.LogOutTime = time.Now().Local()
	user.LoginTime = time.Now().Local()
	user.HeartbeatTime = time.Now().Local()
	db.Create(user)
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
