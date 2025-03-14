package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	HealthCheckRoutes(server)
	AuthRoutes(server, db)
}
