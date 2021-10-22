package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "T2T-GO",
		})
	})

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
