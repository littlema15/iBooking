package models

import "time"

type Booking struct {
	ID        int64 `gorm:"primaryKey" json:"id"`
	UserID    int64
	SeatID    int64
	CreateAt  time.Time
	ExpiredAt time.Time
	IsSigned  bool
}

func (b *Booking) Create() error {
	db.NewRecord(b)
	return db.Create(b).Error
}

func GetBookingByID(id int64) (*Booking, error) {
	var booking Booking
	if err := db.Model(&Booking{}).Where("id =?", id).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func GetBookingByUserID(id int64) ([]Booking, error) {
	var bookings []Booking
	if err := db.Model(&Booking{}).Where("user_id =?", id).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func DeleteBooking(id int64) error {
	return db.Model(&Booking{}).Where("id = ?", id).Delete(&Booking{}).Error
}

func UpdateBooking(id int64, booking *Booking) error {
	return db.Model(&Booking{}).Where("id =?", id).Updates(*booking).Error
}
