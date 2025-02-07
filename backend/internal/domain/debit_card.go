package domain

type DebitCardWithDetailsResponse struct {
	CardName    *string `json:"name"`
	DesignColor *string `json:"color"`
	BorderColor *string `json:"border_color"`
	CardIssuer  *string `json:"issuer"`
	CardNumber  *string `json:"number"`
	CardStatus  *string `json:"status"`
}

type UserDebitCardAggregateResponse struct {
	Cards []DebitCardWithDetailsResponse `json:"cards"`
}
