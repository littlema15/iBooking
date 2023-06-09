package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserInfo struct {
	UserID             int64  `gorm:"primaryKey" json:"user_id"`
	Username           string `gorm:"unique" json:"username"`
	Email              string `json:"email"`
	Gender             string `json:"gender"`
	NumberDefaults     int32  `json:"number_defaults"`
	AcceptNotification bool   `json:"accept_notification"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (u *UserInfo) Create() error {
	db.NewRecord(u)
	return db.Create(u).Error
}

func GetUserinfoByUserID(id int64) (*UserInfo, *gorm.DB, error) {
	var userinfo UserInfo
	db := db.Model(&UserInfo{}).Where("user_id =?", id).First(&userinfo)
	if err := db.Error; err != nil {
		return nil, nil, err
	}
	return &userinfo, db, nil
}

func GetUserinfoByUsername(username string) (*UserInfo, error) {
	var userinfo UserInfo
	if err := db.Model(&UserInfo{}).Where("username =?", username).First(&userinfo).Error; err != nil {
		return nil, err
	}
	return &userinfo, nil
}

func DeleteUserInfo(id int64) error {
	return db.Model(&UserInfo{}).Where("user_id =?", id).Delete(&UserInfo{}).Error
}
