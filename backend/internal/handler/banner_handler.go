package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	bannerUseCase usecase.BannerUseCase
}

func NewBannerHandler(usecase usecase.BannerUseCase) *BannerHandler {
	return &BannerHandler{usecase}
}

// GetBannerByID godoc
// @Summary      get banner from user id
// @Description  get string by ID
// @Tags         Banner
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  []domain.Banner
// @Router       /api/banner/{id} [get]
func (h *BannerHandler) GetBannerByID(c *gin.Context) {
	userID := c.Param("id")
	banner, err := h.bannerUseCase.GetBannerByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}
	c.JSON(http.StatusOK, banner)
}
