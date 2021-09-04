package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		api.GET("/todos", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user_id":    1,
				"todo_id":    1,
				"todo_title": "Do homework",
			})
		})
	}

	router.Run(":5000")
}
