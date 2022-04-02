package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoListGet() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"name": "Ken",
		})
	}
}
