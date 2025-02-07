package module

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterDebitCardModule(router *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	cardRepo := repository.NewDebitCardRepository(db)
	cardUseCase := usecase.NewDebitCardUseCase(cardRepo)
	cardHandler := handler.NewDebitCardHandler(cardUseCase)

	router.GET("/debit-card/:id", cardHandler.GetDebitCardByID)
}
