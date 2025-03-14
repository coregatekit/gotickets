package routes

import (
	"github.com/coregate/tickets-app/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(server *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
	}
}
