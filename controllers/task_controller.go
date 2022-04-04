package controllers

import (
	"net/http"

	"go-gin-todolist/models"

	"github.com/gin-gonic/gin"
)

// Get
func GetTodoList() gin.HandlerFunc {
	return func(c *gin.Context) {
		todolist := models.TodoList{}

		c.JSON(http.StatusOK, gin.H{
			"results": todolist,
		})
	}
}

// Post
func PostTodoList() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoItem := c.PostForm("todoItem")
		done := c.PostForm("done")

		c.JSON(http.StatusOK, gin.H{
			"results":  "Ok",
			"todoItem": todoItem,
			"done":     done,
		})
	}
}

// PUT

// PATCH

// Delete
