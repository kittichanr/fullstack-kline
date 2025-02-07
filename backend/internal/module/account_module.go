package module

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterAccountModule(router *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	userRepo := repository.NewAccountRepository(db)
	accountUseCase := usecase.AccountUseCase(userRepo)
	accountHandler := handler.NewAccountHandler(accountUseCase)

	router.GET("/account/:id", accountHandler.GetAccountByID)
}
