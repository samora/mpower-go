package mpower

type DirectPayPayload struct {
	AccountAlias string  `json:"account_alias"`
	Amount       float64 `json:"amount"`
}

type DirectMobileChargePayload struct {
	CustomerName   string  `json:"customer_name"`
	CustomerPhone  string  `json:"customer_phone"`
	CustomerEmail  string  `json:"customer_email"`
	WalletProvider string  `json:"wallet_provider"`
	MerchantName   string  `json:"merchant_name"`
	Amount         float64 `json:"amount"`
}

type DirectMobileStatusPayload struct {
	Token string `json:"token"`
}

type DirectCardPayload struct {
	Name        string  `json:"card_name"`
	Number      string  `json:"card_number"`
	CVC         string  `json:"cvc"`
	ExpiryMonth string  `json:"exp_month"`
	ExpiryYear  string  `json:"exp_year"`
	Amount      float64 `json:"amount"`
}
