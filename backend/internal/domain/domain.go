package domain

import (
	"time"
)

type AccountBalance struct {
	AccountID string   `gorm:"primaryKey" json:"account_id"`
	UserID    *string  `json:"user_id"`
	Amount    *float64 `json:"amount"`
	DummyCol4 *string  `json:"dummy_col_4"`
}

type AccountDetail struct {
	AccountID     string  `gorm:"primaryKey" json:"account_id"`
	UserID        *string `json:"user_id"`
	Color         *string `json:"color"`
	IsMainAccount *bool   `json:"is_main_account"`
	Progress      *int    `json:"progress"`
	DummyCol5     *string `json:"dummy_col_5"`
}

type AccountFlag struct {
	FlagID    uint      `gorm:"primaryKey;autoIncrement" json:"flag_id"`
	AccountID string    `json:"account_id"`
	UserID    string    `json:"user_id"`
	FlagType  string    `json:"flag_type"`
	FlagValue string    `json:"flag_value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Account struct {
	AccountID     string  `gorm:"primaryKey" json:"account_id"`
	UserID        *string `json:"user_id"`
	Type          *string `json:"type"`
	Currency      *string `json:"currency"`
	AccountNumber *string `json:"account_number"`
	Issuer        *string `json:"issuer"`
	DummyCol3     *string `json:"dummy_col_3"`
}

type Banner struct {
	BannerID    string  `gorm:"primaryKey" json:"banner_id"`
	UserID      *string `json:"user_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	DummyCol11  *string `json:"dummy_col_11"`
}

type DebitCardDesign struct {
	CardID      string  `gorm:"primaryKey" json:"card_id"`
	UserID      *string `json:"user_id"`
	Color       *string `json:"color"`
	BorderColor *string `json:"border_color"`
	DummyCol9   *string `json:"dummy_col_9"`
}

func (DebitCardDesign) TableName() string {
	return "debit_card_design"
}

type DebitCardDetail struct {
	CardID     string  `gorm:"primaryKey" json:"card_id"`
	UserID     *string `json:"user_id"`
	Issuer     *string `json:"issuer"`
	Number     *string `json:"number"`
	DummyCol10 *string `json:"dummy_col_10"`
}

type DebitCardStatus struct {
	CardID    string  `gorm:"primaryKey" json:"card_id"`
	UserID    *string `json:"user_id"`
	Status    *string `json:"status"`
	DummyCol8 *string `json:"dummy_col_8"`
}

func (DebitCardStatus) TableName() string {
	return "debit_card_status"
}

type DebitCard struct {
	CardID    string  `gorm:"primaryKey" json:"card_id"`
	UserID    *string `json:"user_id"`
	Name      *string `json:"name"`
	DummyCol7 *string `json:"dummy_col_7"`
}

type Transaction struct {
	TransactionID string  `gorm:"primaryKey" json:"transaction_id"`
	UserID        *string `json:"user_id"`
	Name          *string `json:"name"`
	Image         *string `json:"image"`
	IsBank        *bool   `json:"is_bank"`
	DummyCol6     *string `json:"dummy_col_6"`
}

type UserGreeting struct {
	UserID    string  `gorm:"primaryKey" json:"user_id"`
	Greeting  *string `json:"greeting"`
	DummyCol2 *string `json:"dummy_col_2"`
}

type User struct {
	UserID    string  `gorm:"primaryKey" json:"user_id"`
	Name      *string `json:"name"`
	DummyCol1 *string `json:"dummy_col_1"`
	PinHash   string  `json:"pin_hash"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
