package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/controllers/middlewares"
	"github.com/littlema15/iBooking/pkg/models"
	"github.com/littlema15/iBooking/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateAdmin(c *gin.Context) {
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if json["username"] == nil || json["password"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is nil",
		})
		return
	}
	// encrypt the password, using password + salt(a string of random numbers) and then hash
	hash, err := bcrypt.GenerateFromPassword([]byte(json["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	admin := models.Administrator{
		ID:       utils.GetID(),
		Username: json["username"].(string),
		Password: string(hash),
	}

	// log.Printf("password: %v\n", string(hash))

	if err := admin.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create admin success",
		"data":    admin,
	})
}

var AdminAuthMiddleware, adminGinJWTMiddleErr = middlewares.GenerateAdminAuthMiddleware()

func AdminLogin(c *gin.Context) {
	if adminGinJWTMiddleErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": adminGinJWTMiddleErr.Error(),
		})
	}
	AdminAuthMiddleware.LoginHandler(c)
}
