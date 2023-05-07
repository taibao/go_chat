package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name       string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Password   string `gorm:"column:password" db:"password" json:"password" form:"password"`
	Phone      string `valid:"matches(^1[3-9]{1}\\d{9}$)" gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
	Email      string `valid:"email" gorm:"column:email" db:"email" json:"email" form:"email"`
	Identity   string `gorm:"column:identity" db:"identity" json:"identity" form:"identity"`
	Salt       string `gorm:"column:salt"`
	ClientIp   string `gorm:"column:client_ip" db:"client_ip" json:"client_ip" form:"client_ip"`
	ClientPort string `gorm:"column:client_port" db:"client_port" json:"client_port" form:"client_port"`
	IsLogout   int64  `gorm:"column:is_logout" db:"is_logout" json:"is_logout" form:"is_logout"`
	DeviceInfo string `gorm:"column:device_info" db:"device_info" json:"device_info" form:"device_info"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
