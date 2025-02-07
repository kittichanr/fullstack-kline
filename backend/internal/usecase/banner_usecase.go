package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type BannerUseCase interface {
	GetBannerByID(bannerID string) (*[]domain.Banner, error)
}

type bannerUseCase struct {
	bannerRepo repository.BannerRepository
}

func NewBannerUseCase(repo repository.BannerRepository) BannerUseCase {
	return &bannerUseCase{repo}
}

func (u *bannerUseCase) GetBannerByID(userID string) (*[]domain.Banner, error) {
	return u.bannerRepo.GetBannerByID(userID)
}
