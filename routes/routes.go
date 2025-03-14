package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	HealthCheckRoutes(server)
	AuthRoutes(server)
}
