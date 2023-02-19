package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/models"

	gin "github.com/gin-gonic/gin"
)

func init() {

	models.ConnectDataBase()
}

func main() {

	router := gin.Default()

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ok"})
	})

	api := router.Group("/api")
	{
		api.POST("/user/register", controllers.Register)
		api.GET("/user", middlewares.JwtAuthMiddleware(), controllers.CurrentUser)
		api.POST("/user/login", controllers.Login)
	}

	authApi := router.Group("/api/auth")
	{
		authApi.Use(middlewares.JwtAuthMiddleware())
		authApi.GET("/is-authenticated", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "yes"})
		})
	}

	router.Run(":8081") // listen and serve on 0.0.0.0:8081
}
