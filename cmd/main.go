package main

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/bikefit-calculator/internal/controller"
)

func main() {
	server := gin.Default()

	server.POST("/calculation", controller.Calculate)

	server.Run(":9090")
}
