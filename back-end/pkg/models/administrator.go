package models

import (
	"github.com/jinzhu/gorm"
	"github.com/littlema15/iBooking/pkg/config"
	"time"
)

var db *gorm.DB

type Administrator struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Administrator{})
	db.AutoMigrate(&Booking{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Seat{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserInfo{})
}

func (a *Administrator) Create() error {
	db.NewRecord(&a)
	return db.Create(a).Error
}

func GetAdminByUsername(username string) (*Administrator, error) {
	var admin Administrator
	if err := db.Model(&Administrator{}).Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
