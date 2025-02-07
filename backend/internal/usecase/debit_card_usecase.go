package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type DebitCardUseCase interface {
	GetDebitCardByID(userID string) (*domain.UserDebitCardAggregateResponse, error)
}

type debitCardUseCase struct {
	cardRepo repository.DebitCardRepository
}

func NewDebitCardUseCase(repo repository.DebitCardRepository) DebitCardUseCase {
	return &debitCardUseCase{repo}
}

func (u *debitCardUseCase) GetDebitCardByID(userID string) (*domain.UserDebitCardAggregateResponse, error) {
	return u.cardRepo.GetDebitCardByID(userID)
}
