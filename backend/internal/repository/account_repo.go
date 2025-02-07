package repository

import (
	"backend/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByID(userID string) (*domain.UserAccountAggregateResponse, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) GetAccountByID(userID string) (*domain.UserAccountAggregateResponse, error) {
	var accounts []domain.Account

	if err := r.db.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		return nil, errors.New("no accounts found for user")
	}

	var aggregatedAccounts []domain.AccountWithDetailsResponse
	for _, account := range accounts {
		var balance domain.AccountBalance
		var detail domain.AccountDetail
		var flags []domain.AccountFlag

		r.db.Where("account_id = ?", account.AccountID).First(&balance)
		r.db.Where("account_id = ?", account.AccountID).First(&detail)
		r.db.Where("account_id = ?", account.AccountID).Find(&flags)

		var flagResponses []domain.AccountFlagResponse
		for _, flag := range flags {
			flagResponses = append(flagResponses, domain.AccountFlagResponse{
				FlagID:    flag.FlagID,
				FlagType:  flag.FlagType,
				FlagValue: flag.FlagValue,
			})
		}

		accountWithDetails := domain.AccountWithDetailsResponse{
			Account: domain.AccountResponse{
				Type:          account.Type,
				Currency:      account.Currency,
				AccountNumber: account.AccountNumber,
				Issuer:        account.Issuer,
				DummyCol3:     account.DummyCol3,
			},
			Flags: flagResponses,
		}

		if balance.AccountID != "" {
			accountWithDetails.Balance = &domain.AccountBalanceResponse{
				Amount: balance.Amount,
			}
		}

		if detail.AccountID != "" {
			accountWithDetails.Detail = &domain.AccountDetailResponse{
				Color:         detail.Color,
				IsMainAccount: detail.IsMainAccount,
				Progress:      detail.Progress,
			}
		}

		aggregatedAccounts = append(aggregatedAccounts, accountWithDetails)
	}

	aggregate := &domain.UserAccountAggregateResponse{
		Accounts: aggregatedAccounts,
	}

	return aggregate, nil
}
