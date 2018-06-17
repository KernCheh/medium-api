package main

import (
	"github.com/Kerncheh/medium-api/controllers/posts"
	"github.com/Kerncheh/medium-api/db/connection"
	"github.com/gin-gonic/gin"
)

func main() {
	connection.Connect()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1/posts")
	{
		v1.GET("/", posts.Index)
		v1.POST("/", posts.Create)

		v1.PUT("/:id", posts.Update)
		v1.PATCH("/:id", posts.Update)
		v1.POST("/:id", posts.Update)

		v1.GET("/:id", posts.Show)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}
