package handlers

import (
	"net/http"

	"github.com/coregate/tickets-app/common"
	"github.com/coregate/tickets-app/pkg/auth"
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
	var createUser auth.CreateUser
	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.authService.Register(createUser); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code:    http.StatusOK,
		Message: "User registered successfully",
		Data:    nil,
	})
}

// GoTemplate 	godoc
// @Summary		Login
// @Description	Login user
// @ID 			Login
// @Tags        Auth
// @Accept		json
// @Produce		json
// @Param		body body auth.LoginRequest true "User data"
// @Success 	200 {object} common.Response "OK"
// @Failure		400 {object} common.Response "Bad Request"
// @Failure		500 {object} common.Response "Internal Server Error"
// @Router			/api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest auth.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	result, err := h.authService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code:    http.StatusOK,
		Message: "User logged in successfully",
		Data:    result,
	})

}
