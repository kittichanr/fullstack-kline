package module

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterUserModule(router *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo, cfg.JwtSecret)
	userHandler := handler.NewUserHandler(userUseCase)

	router.GET("/user/greeting/:id", userHandler.GetGreetingByUserID)
}
