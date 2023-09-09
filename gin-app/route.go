package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rongan/otel-gin-instrument/gin-app/controllers"
	"github.com/rongan/otel-gin-instrument/gin-app/log"
	"net/http"
)

func installRouters(g *gin.Engine) error {
	// 注册 /healthz handler.
	g.GET("/ping", func(c *gin.Context) {
		log.C(c).Infow("ping function called")
		c.JSONP(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Routes
	g.GET("/books", controllers.FindBooks)
	g.GET("/books/:id", controllers.FindBook)
	g.POST("/books", controllers.CreateBook)
	g.PATCH("/books/:id", controllers.UpdateBook)
	g.DELETE("/books/:id", controllers.DeleteBook)
	g.POST("/user/login", controllers.UserLogin)
	g.GET("/user/info", controllers.UserInfo)
	g.POST("/user/logout", controllers.UserLogout)

	return nil
}
