package routes

import (
	"github.com/coregate/tickets-app/handlers"
	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(server *gin.Engine) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	server.GET("/api/health", healthCheckHandler.HealthCheck)
}
