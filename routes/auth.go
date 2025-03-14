package routes

import (
	repositoriese "github.com/coregate/tickets-app/database/repositories"
	"github.com/coregate/tickets-app/handlers"
	"github.com/coregate/tickets-app/packages/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(server *gin.Engine, db *gorm.DB) {
	userRepo := repositoriese.NewUsersRepository(db)

	authService := auth.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
	}
}
