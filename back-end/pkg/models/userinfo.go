package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfo struct {
	UserID             int64  `gorm:"primaryKey" json:"user_id"`
	Username           string `gorm:"primaryKey" json:"username"`
	Email              string `json:"email"`
	Gender             string `json:"gender"`
	NumberDefaults     int32  `json:"number_defaults"`
	AcceptNotification bool   `json:"accept_notification"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func GetUserinfoByUserID(id int64) (*UserInfo, *gorm.DB, error) {
	var user UserInfo
	err := db.Model(&UserInfo{}).Where("userid =?", id).First(&user).Error
	return &user, db, err
}

func GetUserinfoByUsername(username string) (*UserInfo, *gorm.DB, error) {
	var user UserInfo
	err := db.Model(&UserInfo{}).Where("username =?", username).First(&user).Error
	return &user, db, err
}

func DeleteUserInfo(id int64) error {
	var user UserInfo
	return db.Model(&UserInfo{}).Where("id =?", id).Delete(&user).Error
}
