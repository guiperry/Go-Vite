package api

import (
	"backend/internal/api/handlers"
	"backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	v1.Use(middleware.Logger())
	{
		v1.GET("/items", handlers.ListItems)
		v1.POST("/items", handlers.CreateItem)
		v1.GET("/items/:id", handlers.GetItem)
		v1.PUT("/items/:id", handlers.UpdateItem)
		v1.DELETE("/items/:id", handlers.DeleteItem)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
