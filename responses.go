package mpower

type DirectMobileChargeResponse struct {
	ResponseCode        string `json:"response_code"`
	ResponseText        string `json:"response_text"`
	Description         string `json:"description"`
	TransactionID       string `json:"transaction_id"`
	Token               string `json:"token"`
	MobileInvoiceNumber string `json:"mobile_invoice_number"`
}
