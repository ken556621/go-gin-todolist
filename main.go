package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("KKKKKKKKKKKKKKKK============")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "KKKKKKKKKKKKKKKK============",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
