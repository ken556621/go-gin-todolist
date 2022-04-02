package routes

import (
	"go-gin-todolist/controllers"

	"github.com/gin-gonic/gin"
)

func TodoListRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/signup", controllers.TodoListGet())
}
