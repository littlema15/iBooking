package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/littlema15/iBooking/pkg/models"
	"net/http"
)

// Notify TODO: complete email notification
func Notify(c *gin.Context) {
	if c.Param("username") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username cannot be empty",
		})
	}
	userinfo, err := models.GetUserinfoByUsername(c.Param("username"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	emailAddress := userinfo.Email
	if emailAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user has no email address",
		})
		return
	}
	// send notification by email
	e := email.NewEmail()
	e.To = []string{emailAddress}
	e.Subject = "iBooking"
	e.Text = []byte("You reservation will due in xxx time")
}
