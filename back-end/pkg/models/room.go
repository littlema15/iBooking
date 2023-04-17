package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Room struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	RoomNumber string `json:"room_number"`
	Location   string `json:"location"`
	Seats      []Seat `json:"seats"`
	Total      int16  `json:"total"`
	Free       int16  `json:"free"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *Room) Create() error {
	db.NewRecord(r)
	return db.Create(r).Error
}

func GetAllRooms() ([]Room, error) {
	var rooms []Room
	if err := db.Model(&Room{}).Preload("Seats").Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func GetRoomById(id int64) (*Room, *gorm.DB, error) {
	var room Room
	if err := db.Model(&Room{}).Where("ID=?", id).Preload("Seats").Find(&room).Error; err != nil {
		return nil, nil, err
	}
	return &room, db, nil
}

func DeleteRoom(id int64) error {
	return db.Model(&Room{}).Where("ID=?", id).Delete(&Room{}).Error
}
