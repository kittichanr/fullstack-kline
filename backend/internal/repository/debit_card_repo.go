package repository

import (
	"backend/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type DebitCardRepository interface {
	GetDebitCardByID(userID string) (*domain.UserDebitCardAggregateResponse, error)
}

type debitCardRepository struct {
	db *gorm.DB
}

func NewDebitCardRepository(db *gorm.DB) DebitCardRepository {
	return &debitCardRepository{db}
}

func (r *debitCardRepository) GetDebitCardByID(userID string) (*domain.UserDebitCardAggregateResponse, error) {
	var cards []domain.DebitCard

	if err := r.db.Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return nil, errors.New("no debit cards found for user")
	}

	var aggregatedCards []domain.DebitCardWithDetailsResponse
	for _, card := range cards {
		var design domain.DebitCardDesign
		var detail domain.DebitCardDetail
		var status domain.DebitCardStatus

		r.db.Where("card_id = ?", card.CardID).First(&design)
		r.db.Where("card_id = ?", card.CardID).First(&detail)
		r.db.Where("card_id = ?", card.CardID).First(&status)

		cardWithDetails := domain.DebitCardWithDetailsResponse{
			CardName:    card.Name,
			DesignColor: design.Color,
			BorderColor: design.BorderColor,
			CardIssuer:  detail.Issuer,
			CardNumber:  detail.Number,
			CardStatus:  status.Status,
		}

		aggregatedCards = append(aggregatedCards, cardWithDetails)
	}

	aggregate := &domain.UserDebitCardAggregateResponse{
		Cards: aggregatedCards,
	}

	return aggregate, nil
}
