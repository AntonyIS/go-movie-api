package main

import (
	"github.com/AntonyIS/GoProject/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Defualt()

	router.GET("/api/v1", controller.Home)
	router.Run("localhost:5000")
}
