package module

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterAuthModule(router *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	userRepo := repository.NewUserRepository(db)
	authUseCase := usecase.NewAuthUseCase(userRepo, cfg.JwtSecret)
	authHandler := handler.NewAuthHandler(authUseCase)

	router.POST("/set-pin", authHandler.SetPin)
	router.POST("/auth-pin", authHandler.AuthenticatePin)
}
