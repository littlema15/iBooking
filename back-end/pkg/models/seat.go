package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Seat struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	RoomID    int64 `json:"room_id"`
	X         int8  `json:"x"`
	Y         int8  `json:"y"`
	Status    int8  `json:"status"`
	Plug      bool  `json:"plug"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Seat) Create() error {
	db.NewRecord(s)
	return db.Create(s).Error
}

func GetAllSeats() ([]Seat, error) {
	var seats []Seat
	if err := db.Model(&Seat{}).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func GetSeatByID(id int64) (*Seat, *gorm.DB, error) {
	var seat Seat
	db := db.Model(&Seat{}).Where("ID=?", id).First(&seat)
	if err := db.Error; err != nil {
		return nil, nil, err
	}
	return &seat, db, nil
}

func DeleteSeat(id int64) error {
	var seat Seat
	if err := db.Model(&Seat{}).Where("ID=?", id).Delete(&seat).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSeatByRoomID(id int64) error {
	return db.Model(&Seat{}).Where("room_id=?", id).Delete(&Seat{}).Error
}
