package main

import (
	"gin-bun-todos/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome home",
		})
	})
	routers.RoutesHandler(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
