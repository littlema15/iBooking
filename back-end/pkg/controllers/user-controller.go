package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/littlema15/iBooking/pkg/controllers/middlewares"
	"github.com/littlema15/iBooking/pkg/models"
	"github.com/littlema15/iBooking/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// CreateUser TODO:test
func CreateUser(c *gin.Context) {
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
	user := models.User{
		ID:       utils.GetID(),
		Username: json["username"].(string),
		Password: string(hash),
	}

	// log.Printf("password: %v\n", string(hash))
	// TODO:test
	userinfo := models.UserInfo{}
	if json["userinfo"] != nil {
		info := json["userinfo"].(map[string]interface{})
		if info["email"] != nil {
			userinfo.Email = info["email"].(string)
		}
		if info["gender"] != nil {
			userinfo.Gender = info["gender"].(string)
		}
		if info["number_defaults"] != nil {
			userinfo.NumberDefaults = utils.Stoi(info["number_defaults"].(string), 32).(int32)
		}
		if info["accept_notification"] != nil {
			userinfo.AcceptNotification = info["accept_notification"].(bool)
		}
		log.Printf("userinfo: %v\n", userinfo)
	}

	if err := user.Create(userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
		"data":    user,
	})
}

var UserAuthMiddleware, userGinJWTMiddleErr = middlewares.GenerateUserAuthMiddleware()

func UserLogin(c *gin.Context) {
	if userGinJWTMiddleErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": userGinJWTMiddleErr.Error(),
		})
		return
	}
	UserAuthMiddleware.LoginHandler(c)
}

func UserLogout(c *gin.Context) {
	if userGinJWTMiddleErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": userGinJWTMiddleErr.Error(),
		})
		return
	}
	UserAuthMiddleware.LogoutHandler(c)
}

func UserRefreshToken(c *gin.Context) {
	if userGinJWTMiddleErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": userGinJWTMiddleErr.Error(),
		})
		return
	}
	UserAuthMiddleware.RefreshToken(c)
}

// DeleteUser TODO:test
func DeleteUser(c *gin.Context) {
	if c.Param("userID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userID is required",
		})
		return
	}
	userID := utils.Stoi(c.Param("userID"), 64).(int64)
	if err := models.DeleteUserInfo(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user OK",
	})
}

func UpdateUser(c *gin.Context) {
	if c.Param("userID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userID is required",
		})
		return
	}
	userID := utils.Stoi(c.Param("userID"), 64).(int64)
	json := make(map[string]interface{})
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userinfo, db, err := models.GetUserinfoByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if json["email"] != nil {
		userinfo.Email = json["email"].(string)
	}
	if json["gender"] != nil {
		userinfo.Gender = json["gender"].(string)
	}
	if json["number_defaults"] != nil {
		userinfo.NumberDefaults = utils.Stoi(json["number_defaults"].(string), 32).(int32)
	}
	if json["accept_notification"] != nil {
		userinfo.AcceptNotification = json["accept_notification"].(bool)
	}
	if err := db.Save(userinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update user OK",
		"data":    userinfo,
	})
}

func GetUserByID(c *gin.Context) {
	if c.Param("userID") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userID is required",
		})
		return
	}
	userID := utils.Stoi(c.Param("userID"), 64).(int64)
	user, _, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get user OK",
		"data":    user,
	})
}
