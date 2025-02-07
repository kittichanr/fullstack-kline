package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	useCase usecase.AccountUseCase
}

func NewAccountHandler(useCase usecase.AccountUseCase) *AccountHandler {
	return &AccountHandler{useCase}
}

// GetAccountByID godoc
// @Summary      get account from user id
// @Description  get string by ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  domain.UserAccountAggregateResponse
// @Router       /api/accounts/{id} [get]
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id := c.Param("id")
	account, err := h.useCase.GetAccountByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}
