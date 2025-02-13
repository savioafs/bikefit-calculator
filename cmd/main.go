package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.POST("/calculation")

	server.Run(":9090")
}
