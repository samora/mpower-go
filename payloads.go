package mpower

type directPayPayload struct {
	AccountAlias string  `json:"account_alias"`
	Amount       float64 `json:"amount"`
}

type directMobileChargePayload struct {
	CustomerName   string  `json:"customer_name"`
	CustomerPhone  string  `json:"customer_phone"`
	CustomerEmail  string  `json:"customer_email"`
	WalletProvider string  `json:"wallet_provider"`
	MerchantName   string  `json:"merchant_name"`
	Amount         float64 `json:"amount"`
}

type directMobileStatusPayload struct {
	Token string `json:"token"`
}
