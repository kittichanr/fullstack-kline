package domain

type AccountResponse struct {
	Type          *string `json:"type"`
	Currency      *string `json:"currency"`
	AccountNumber *string `json:"account_number"`
	Issuer        *string `json:"issuer"`
	DummyCol3     *string `json:"dummy_col_3"`
}

type AccountBalanceResponse struct {
	Amount *float64 `json:"amount,omitempty"`
}

type AccountDetailResponse struct {
	Color         *string `json:"color,omitempty"`
	IsMainAccount *bool   `json:"is_main_account,omitempty"`
	Progress      *int    `json:"progress,omitempty"`
}

type AccountFlagResponse struct {
	FlagID    uint   `json:"flag_id"`
	FlagType  string `json:"flag_type"`
	FlagValue string `json:"flag_value"`
}

type AccountWithDetailsResponse struct {
	Account AccountResponse         `json:"account"`
	Balance *AccountBalanceResponse `json:"balance,omitempty"`
	Detail  *AccountDetailResponse  `json:"detail,omitempty"`
	Flags   []AccountFlagResponse   `json:"flags,omitempty"`
}

type UserAccountAggregateResponse struct {
	Accounts []AccountWithDetailsResponse `json:"accounts"`
}
