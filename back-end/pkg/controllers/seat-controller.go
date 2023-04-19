package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/models"
	"github.com/littlema15/iBooking/pkg/utils"
	"net/http"
)

func CreateSeat(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var errorMsg bytes.Buffer

	var Seats []models.Seat
	for _, v := range json {
		m := v.(map[string]interface{})
		// check if the room exists
		roomID := utils.Stoi(m["room_id"].(string), 64).(int64)
		if _, err := models.GetRoomById(roomID); err != nil {
			errorMsg.WriteString(err.Error() + "\n")
			continue
		}
		seat := &models.Seat{
			ID:     utils.GetID(),
			RoomID: roomID,
			X:      utils.Stoi(m["x"].(string), 8).(int8),
			Y:      utils.Stoi(m["y"].(string), 8).(int8),
			Status: utils.Stoi(m["status"].(string), 8).(int8),
			Plug:   m["plug"].(bool),
		}
		if err := seat.Create(); err != nil {
			errorMsg.WriteString(err.Error() + "\n")
			continue
		}
		Seats = append(Seats, *seat)
	}
	if errorMsg.Len() > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMsg.String(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"create seats": Seats,
	})
}

func UpdateSeat(c *gin.Context) {
	json := make(map[string]map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// log.Printf("%v\n",&json)
	var errorMsg bytes.Buffer

	var Seats []models.Seat

	for _, v := range json {
		seatID := utils.Stoi(v["id"].(string), 64).(int64)
		seat, err := models.GetSeatByID(seatID)
		if err != nil {
			errorMsg.WriteString(err.Error() + "\n")
			continue
		}
		if err := models.DeleteSeat(seat.ID); err != nil {
			errorMsg.WriteString(err.Error() + "\n")
			continue
		}
		if v["room_id"] != nil {
			roomID := utils.Stoi(v["room_id"].(string), 64).(int64)
			if _, err := models.GetRoomById(roomID); err != nil {
				errorMsg.WriteString(err.Error() + "\n")
				continue
			}
			seat.RoomID = roomID
		}
		if v["x"] != nil {
			seat.X = utils.Stoi(v["x"].(string), 8).(int8)
		}
		if v["y"] != nil {
			seat.Y = utils.Stoi(v["y"].(string), 8).(int8)
		}
		if v["status"] != nil {
			seat.Status = utils.Stoi(v["status"].(string), 8).(int8)
		}
		if v["plug"] != nil {
			seat.Plug = v["plug"].(bool)
		}

		if err := seat.Create(); err != nil {
			errorMsg.WriteString(err.Error() + "\n")
			continue
		}
		Seats = append(Seats, *seat)
	}

	if errorMsg.Len() > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMsg.String(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"seats": Seats,
	})
}

// GetSeat return all the seats
func GetSeat(c *gin.Context) {
	seats, err := models.GetAllSeats()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"seats": seats,
	})
}

func DeleteSeat(c *gin.Context) {
	if c.Param("seatID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "seatID is required",
		})
		return
	}
	seatID := utils.Stoi(c.Param("seatID"), 64).(int64)
	if err := models.DeleteSeat(seatID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete seat OK",
	})
}

func GetSeatByID(c *gin.Context) {
	if c.Param("seatID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "seatID is required",
		})
		return
	}
	seatID := utils.Stoi(c.Param("seatID"), 64).(int64)
	seat, err := models.GetSeatByID(seatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"seat": seat,
	})
}
