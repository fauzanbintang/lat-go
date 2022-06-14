package routers

import (
	"github.com/gin-gonic/gin"
)

func RoutesHandler(r *gin.Engine) {
	todoHandler := NewTodoHandler()

	todos := r.Group("/todos")
	{
		// todos.GET("/", getAllTodos)
		// todos.GET("/:tid", getAllTodos)
		// todos.POST("/", getAllTodos)
		// todos.PUT("/", getAllTodos)
		// todos.DELETE("/", getAllTodos)
	}
}
