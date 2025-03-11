package routes

import "github.com/gin-gonic/gin"

func HealthCheckRoutes(server *gin.Engine) {
	server.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "I'm fine. 괜찮아 🚀",
		})
	})
}
