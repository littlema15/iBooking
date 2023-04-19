package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/models"
	"github.com/littlema15/iBooking/pkg/utils"
	"net/http"
)

func BookSeat(c *gin.Context) {
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json["user_id"] == nil || json["seat_id"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id or seat_id is required",
		})
		return
	}
	var booking = models.Booking{
		ID:       utils.GetID(),
		UserID:   utils.Stoi(json["user_id"].(string), 64).(int64),
		SeatID:   utils.Stoi(json["seat_id"].(string), 64).(int64),
		IsSigned: false,
	}
	if err := booking.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "booking created successfully",
		"data":    booking,
	})
}

func GetBookingByUserID(c *gin.Context) {
	if c.Param("userID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userID is required",
		})
		return
	}
	userID := utils.Stoi(c.Param("userID"), 64).(int64)
	if _, err := models.GetUserByID(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	bookings, err := models.GetBookingByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "bookings retrieved successfully",
		"data":    bookings,
	})
}

func GetBookingByID(c *gin.Context) {
	if c.Param("bookingID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bookingID is required",
		})
		return
	}
	bookingID := utils.Stoi(c.Param("bookingID"), 64).(int64)
	booking, err := models.GetBookingByID(bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "booking retrieved successfully",
		"data":    booking,
	})
}

func DeleteBooking(c *gin.Context) {
	if c.Param("bookingID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bookingID is required",
		})
		return
	}
	bookingID := utils.Stoi(c.Param("bookingID"), 64).(int64)
	if err := models.DeleteBooking(bookingID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "booking deleted successfully",
	})
}

func UpdateBooking(c *gin.Context) {
	if c.Param("bookID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bookingID is required",
		})
		return
	}
	bookingID := utils.Stoi(c.Param("bookingID"), 64).(int64)
	booking, err := models.GetBookingByID(bookingID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	var json map[string]interface{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if json["is_signed"] != nil {
		if booking.IsSigned == json["is_signed"].(bool) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "not changed",
			})
		}
		booking.IsSigned = json["is_signed"].(bool)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "booking updated successfully",
		"data":    booking,
	})
}
