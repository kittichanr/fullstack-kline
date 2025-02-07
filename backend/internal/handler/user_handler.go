package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase}
}

// GetGreetingByUserID godoc
// @Summary      get greeting from user id
// @Description  get string by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  domain.UserGreeting
// @Router       /api/user/greeting/{id} [get]
func (c *UserHandler) GetGreetingByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")
	greeting, err := c.useCase.GetGreetingByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if greeting == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Greeting not found"})
		return
	}
	ctx.JSON(http.StatusOK, greeting)
}
