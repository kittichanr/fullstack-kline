package module

import (
	"backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterBannerModule(router *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	bannerRepo := repository.NewBannerRepository(db)
	bannerUseCase := usecase.NewBannerUseCase(bannerRepo)
	bannerHandler := handler.NewBannerHandler(bannerUseCase)

	router.GET("/banner/:id", bannerHandler.GetBannerByID)
}
