package main

import (
	"github.com/gin-gonic/gin"

	"go-gin-todolist/routes"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	routes.TodoListRoutes(router)

	port := "8080"
	router.Run(":" + port)
}
