package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type AccountUseCase interface {
	GetAccountByID(id string) (*domain.UserAccountAggregateResponse, error)
}

type accountUsecase struct {
	repo repository.AccountRepository
}

func NewAccountUsecase(repo repository.AccountRepository) AccountUseCase {
	return &accountUsecase{repo}
}

func (u *accountUsecase) GetAccountByID(id string) (*domain.UserAccountAggregateResponse, error) {
	return u.repo.GetAccountByID(id)
}
