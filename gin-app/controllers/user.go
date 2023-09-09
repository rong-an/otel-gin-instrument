package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rongan/otel-gin-instrument/gin-app/models"
	"log"
	"net/http"
	"time"
)

func UserInfo(c *gin.Context) {
	userInfo := &models.UserInfo{
		ID:           1,
		Name:         "admin",
		LoginTime:    time.Time{},
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a user.",
		Roles:        []string{"admin"},
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   userInfo,
	})
}

func UserLogin(c *gin.Context) {
	//user := &models.User{
	//	UserName: "admin",
	//	Password: "123456",
	//}
	user := &models.User{}

	c.ShouldBind(&user)
	log.Println(user.UserName, user.Password)
	if user.UserName == "admin" && user.Password == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  0,
			"errmsg": "",
			"data": gin.H{
				"token": "admin",
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"errno":  400,
			"errmsg": "账号密码错误",
			"data": gin.H{
				"token": "admin",
			},
		})
	}
}

func UserLogout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data": gin.H{
			"token": "admin",
		},
	})
}
