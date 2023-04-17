package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/models"
	"github.com/littlema15/iBooking/pkg/utils"
	"net/http"
)

func CreateRoom(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// log.Printf("%v\n", &json)

	room := &models.Room{
		ID:         utils.GetID(),
		RoomNumber: json["room_number"].(string),
		Location:   json["location"].(string),
		Total:      utils.Stoi(json["total"].(string), 16).(int16),
		Free:       utils.Stoi(json["free"].(string), 16).(int16),
	}
	if err := room.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"room": room,
	})
}

func GetRoom(c *gin.Context) {
	rooms, err := models.GetAllRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}

func GetRoomByID(c *gin.Context) {
	if c.Param("roomID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty id"})
	}
	roomID := utils.Stoi(c.Param("roomID"), 64).(int64)
	room, _, err := models.GetRoomById(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"room": room,
	})
}

// DeleteRoom will delete a room from the database with given id
func DeleteRoom(c *gin.Context) {
	if c.Param("roomID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "roomID is required",
		})
	}
	roomID := utils.Stoi(c.Param("roomID"), 64).(int64)
	// delete the seat in the room
	if err := models.DeleteSeatByRoomID(roomID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.DeleteRoom(roomID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete OK",
	})
}

// UpdateRoom delete and create new room for update
func UpdateRoom(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// log.Printf("%v\n", &json)
	if json["room_id"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "roomID is required",
		})
	}
	roomID := utils.Stoi(json["room_id"].(string), 64).(int64)
	room, db, err := models.GetRoomById(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if json["room_number"] != nil {
		room.RoomNumber = json["room_number"].(string)
	}
	if json["location"] != nil {
		room.Location = json["location"].(string)
	}
	// using UpdateSeat to update,not here
	//if json["seat"] != nil {
	//	updateSeats(json["seat"].(map[string]interface{}))
	//}
	if json["total"] != nil {
		room.Total = utils.Stoi(json["total"].(string), 16).(int16)
	}
	if json["free"] != nil {
		room.Free = utils.Stoi(json["free"].(string), 16).(int16)
	}
	if err := db.Save(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updateRoom": room,
	})
}
