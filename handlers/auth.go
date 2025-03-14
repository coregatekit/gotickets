package handlers

import (
	"github.com/coregate/tickets-app/packages/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService auth.IAuthService
}

func NewAuthHandler(authService auth.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// GoTemplate 	godoc
// @Summary		Register
// @Description	Register new user
// @ID 			Register
// @Tags        Auth
// @Accept		json
// @Produce		json
// @Param		body body auth.CreateUser true "User data"
// @Success 	200 {object} common.Response "OK"
// @Failure		400 {object} common.Response "Bad Request"
// @Failure		500 {object} common.Response "Internal Server Error"
// @Router			/api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "User registered",
	})
}
