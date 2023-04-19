package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	DeletedAt     time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
	Name          string    `gorm:"column:name" db:"name" json:"name" form:"name"`
	Password      string    `gorm:"column:password" db:"password" json:"password" form:"password"`
	Phone         string    `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
	Email         string    `gorm:"column:email" db:"email" json:"email" form:"email"`
	Identity      string    `gorm:"column:identity" db:"identity" json:"identity" form:"identity"`
	ClientIp      string    `gorm:"column:client_ip" db:"client_ip" json:"client_ip" form:"client_ip"`
	ClientPort    string    `gorm:"column:client_port" db:"client_port" json:"client_port" form:"client_port"`
	LoginTime     time.Time `gorm:"column:login_time" db:"login_time" json:"login_time" form:"login_time"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time" db:"heartbeat_time" json:"heartbeat_time" form:"heartbeat_time"`
	LogOutTime    time.Time `gorm:"column:log_out_time" db:"log_out_time" json:"log_out_time" form:"log_out_time"`
	IsLogout      int64     `gorm:"column:is_logout" db:"is_logout" json:"is_logout" form:"is_logout"`
	DeviceInfo    string    `gorm:"column:device_info" db:"device_info" json:"device_info" form:"device_info"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}
