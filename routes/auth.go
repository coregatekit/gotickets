package routes

import (
	"github.com/coregate/tickets-app/configs"
	"github.com/coregate/tickets-app/database/repos"
	"github.com/coregate/tickets-app/handlers"
	"github.com/coregate/tickets-app/pkg/auth"
	"github.com/coregate/tickets-app/pkg/encryptions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(server *gin.Engine, db *gorm.DB, configs *configs.Configs) {
	userRepo := repos.NewUsersRepository(db)

	encryptionsService := encryptions.NewEncryptionsService(configs)

	authService := auth.NewAuthService(userRepo, encryptionsService)

	authHandler := handlers.NewAuthHandler(authService)
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
	}
}
