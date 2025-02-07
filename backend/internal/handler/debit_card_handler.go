package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DebitCardHandler struct {
	cardUseCase usecase.DebitCardUseCase
}

func NewDebitCardHandler(usecase usecase.DebitCardUseCase) *DebitCardHandler {
	return &DebitCardHandler{usecase}
}

// GetDebitCardByID godoc
// @Summary      get banner from user id
// @Description  get string by ID
// @Tags         Debit Card
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  domain.UserDebitCardAggregateResponse
// @Router       /api/debit-card/{id} [get]
func (h *DebitCardHandler) GetDebitCardByID(c *gin.Context) {
	userID := c.Param("id")
	card, err := h.cardUseCase.GetDebitCardByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}
