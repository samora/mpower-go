package mpower

type Response struct {
	ResponseCode string `json:"response_code"`
	ResponseText string `json:"response_text"`
}

type DirectMobileChargeResponse struct {
	*Response
	Description         string `json:"description"`
	TransactionID       string `json:"transaction_id"`
	Token               string `json:"token"`
	MobileInvoiceNumber string `json:"mobile_invoice_number"`
}

type DirectMobileStatusResponse struct {
	*Response
	Description         string `json:"description"`
	TransactionStatus   string `json:"tx_status"`
	TransactionID       string `json:"transaction_id"`
	MobileInvoiceNumber string `json:"mobile_invoice_number"`
	CancelReason        string `json:"cancel_reason"`
}
