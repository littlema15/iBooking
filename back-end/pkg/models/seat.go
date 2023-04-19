package models

import (
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

func GetSeatByID(id int64) (*Seat, error) {
	var seat Seat
	if err := db.Model(&Seat{}).Where("id=?", id).First(&seat).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}

func DeleteSeat(id int64) error {
	return db.Model(&Seat{}).Where("id=?", id).Delete(&Seat{}).Error
}

func DeleteSeatByRoomID(id int64) error {
	return db.Model(&Seat{}).Where("room_id=?", id).Delete(&Seat{}).Error
}
