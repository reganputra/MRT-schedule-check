package main

import (
	"github.com/gin-gonic/gin"
	"mrt-schedule-checker/modules/station"
)

func main() {

	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running!",
		})
	})

	api := router.Group("/v1/api")

	station.Intialize(api)

	err := router.Run(":1500")
	if err != nil {
		return
	}
}
