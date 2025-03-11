package handlers

import "github.com/gin-gonic/gin"

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// GoTemplate 	godoc
// @Summary		HealthCheck
// @Description	HealthCheck status is ready of service
// @ID 			HealthCheck
// @Tags        HealthCheck
// @Success 	200 {object} common.Response "OK"
// @Failure		500 {object} common.Response "Internal Server Error"
// @Router			/api/health [get]
func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "I'm fine. 괜찮아 🚀",
	})
}
