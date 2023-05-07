package main

import (
	"go_chat/router"
	"go_chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()
	err := r.Run(":1234")
	if err != nil {
		return
	}
}
