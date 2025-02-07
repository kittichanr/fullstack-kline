package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type BannerRepository interface {
	GetBannerByID(userID string) (*[]domain.Banner, error)
}

type bannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{db}
}

func (r *bannerRepository) CreateBanner(banner *domain.Banner) error {
	return r.db.Create(banner).Error
}

func (r *bannerRepository) GetBannerByID(userID string) (*[]domain.Banner, error) {
	var banner []domain.Banner
	err := r.db.Where("user_id = ?", userID).Find(&banner).Error
	return &banner, err
}
