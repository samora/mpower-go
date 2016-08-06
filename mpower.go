package mpower

import (
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jmcvetta/napping"
)

// MPower object.
type MPower struct {
	baseURL string
	setup   Setup
	store   Store
	session *napping.Session
}

// New MPower object
func New(client *http.Client, setup Setup, store Store) *MPower {
	_, err := govalidator.ValidateStruct(&setup)
	if err != nil {
		log.Panicln(err)
	}
	_, err = govalidator.ValidateStruct(&store)
	if err != nil {
		log.Panicln(err)
	}
	mpower := new(MPower)
	mpower.setup = setup
	mpower.store = store
	if mpower.setup.IsLive {
		mpower.baseURL = "https://app.mpowerpayments.com/api/v1/"
	} else {
		mpower.baseURL = "https://app.mpowerpayments.com/sandbox-api/v1/"
	}

	// setup HTTP session
	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("MP-Master-Key", mpower.setup.MasterKey)
	header.Add("MP-Private-Key", mpower.setup.PrivateKey)
	header.Add("MP-Public-Key", mpower.setup.PublicKey)
	header.Add("MP-Token", mpower.setup.Token)

	// http client
	if client == nil {
		client = http.DefaultClient
	}

	mpower.session = &napping.Session{Header: &header, Client: client}

	return mpower
}

// DirectPay transfers funds to another MPower account.
func (m *MPower) DirectPay(account string, amount float64) (*DirectPayResponse, error) {
	payload := directPayPayload{
		AccountAlias: account,
		Amount:       amount,
	}
	response := DirectPayResponse{}
	_, err := m.session.Post(m.baseURL+"/direct-pay/credit-account", &payload, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DirectMobileCharge charges mobile wallet by pushing a bill prompt to handset.
func (m *MPower) DirectMobileCharge(name, email, mobile, wallet string,
	amount float64) (*DirectMobileChargeResponse, error) {

	payload := directMobileChargePayload{
		CustomerName:   name,
		CustomerPhone:  mobile,
		CustomerEmail:  email,
		WalletProvider: wallet,
		MerchantName:   m.store.Name,
		Amount:         amount,
	}
	response := DirectMobileChargeResponse{}
	_, err := m.session.Post(m.baseURL+"/direct-mobile/charge", &payload, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// DirectMobileStatus checks the status of direct mobile charge.
func (m *MPower) DirectMobileStatus(token string) (*DirectMobileStatusResponse, error) {
	payload := directMobileStatusPayload{
		Token: token,
	}
	response := DirectMobileStatusResponse{}
	_, err := m.session.Post(m.baseURL+"/direct-mobile/status", &payload, &response, nil)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
