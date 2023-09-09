package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rongan/otel-gin-instrument/gin-app/models"
)

func main() {
	r := gin.Default()
	models.InitDB()

	installRouters(r)
	r.Run(":8080")
}
