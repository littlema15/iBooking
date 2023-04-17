package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique" form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create is used to Sign up
func (u *User) Create(info UserInfo) error {
	db.NewRecord(u)
	return db.Save(u).Error
}

func GetAllUser() ([]User, error) {
	var users []User
	if err := db.Model(&users).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id int64) (*User, *gorm.DB, error) {
	var user User
	if err := db.Model(&User{}).Where("ID =?", id).First(&user).Error; err != nil {
		return nil, nil, err
	}
	return &user, db, nil
}

func GetUserByUsername(username string) (*User, *gorm.DB, error) {
	var user User
	if err := db.Model(&User{}).Where("username =?", username).First(&user).Error; err != nil {
		return nil, nil, err
	}
	return &user, db, nil
}

func DeleteUser(id int64) error {
	return db.Model(&User{}).Where("ID = ?", id).Delete(&User{}).Error
}
