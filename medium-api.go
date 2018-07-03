package main

import (
	"path"
	"path/filepath"
	"runtime"

	"medium-api/controllers/posts"
	"medium-api/db/connection"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	connection.Connect()
	router := gin.Default()

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	router.Use(cors.Default())

	router.Use(static.Serve("/", static.LocalFile(path.Join(basepath, "views", "dist"), true)))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1/posts")
	{
		v1.GET("", posts.Index)
		v1.POST("", posts.Create)

		v1.PUT("/:id", posts.Update)
		v1.PATCH("/:id", posts.Update)
		v1.POST("/:id", posts.Update)

		v1.GET("/:id", posts.Show)
	}

	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
