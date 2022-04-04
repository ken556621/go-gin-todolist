package routes

import (
	"go-gin-todolist/controllers"

	"github.com/gin-gonic/gin"
)

func TodoListRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tasks", controllers.GetTodoList())
	incomingRoutes.POST("/task", controllers.PostTodoList())
}
