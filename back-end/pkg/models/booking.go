package models

import "time"

type Booking struct {
	ID          int64 `gorm:"primaryKey" json:"id"`
	UserId      int64
	SeatId      int64
	BookingTime time.Time
	IsSigned    bool
}
