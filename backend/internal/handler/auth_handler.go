package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	useCase usecase.AuthUseCase
}

func NewAuthHandler(useCase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{useCase}
}

type PinRequest struct {
	Name string `json:"name"`
	Pin  string `json:"pin"`
}

// SetPin godoc
//
//	@Summary		Set a PIN for a user
//	@Description	Allows a user to set a PIN for authentication
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		PinRequest	true	"User name and PIN"
//	@Success		200		{object}	string		"message: PIN set successfully"
//	@Router			/api/set-pin [post]
func (h *AuthHandler) SetPin(c *gin.Context) {
	var req PinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	err := h.useCase.SetPin(req.Name, req.Pin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PIN set successfully"})
}

// AuthenticatePin godoc
//
//	@Summary		Authenticate user via PIN
//	@Description	Authenticates a user and returns a JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		PinRequest	true	"User name and PIN"
//	@Success		200		{object}	string
//	@Router			/api/auth-pin [post]
func (h *AuthHandler) AuthenticatePin(c *gin.Context) {
	var req PinRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	token, err := h.useCase.AuthenticatePin(req.Name, req.Pin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
