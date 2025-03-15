package routes

import (
	"github.com/coregate/tickets-app/configs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB, configs *configs.Configs) {
	HealthCheckRoutes(server)
	AuthRoutes(server, db, configs)
}
