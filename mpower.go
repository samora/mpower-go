package mpower

import (
	"log"
	"net/http"

	"github.com/samora/mpower-go/Godeps/_workspace/src/github.com/asaskevich/govalidator"
	"github.com/samora/mpower-go/Godeps/_workspace/src/github.com/jmcvetta/napping"
)

type MPower struct {
	baseURL string
	setup   *Setup
	store   *Store
	session *napping.Session
}

// New MPower object
func New(setup *Setup, store *Store) *MPower {
	_, err := govalidator.ValidateStruct(setup)
	if err != nil {
		log.Panicln(err)
	}
	_, err = govalidator.ValidateStruct(store)
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
	mpower.session = &napping.Session{Header: &header}

	return mpower
}

// DirectMobileCharge charges mobile wallet by pushing a bill prompt to handset.
func (m *MPower) DirectMobileCharge(name, email, mobile, wallet string,
	amount float64) (*DirectMobileChargeResponse, error) {

	payload := struct {
		CustomerName   string `json:"customer_name"`
		CustomerPhone  string `json:"customer_phone"`
		CustomerEmail  string `json:"customer_email"`
		WalletProvider string `json:"wallet_provider"`
		MerchantName   string `json:"merchant_name"`
		Amount         string `json:"amount"`
	}{
		CustomerName:   name,
		CustomerPhone:  mobile,
		CustomerEmail:  email,
		WalletProvider: wallet,
		MerchantName:   m.store.Name,
		Amount:         govalidator.ToString(amount),
	}
	response := new(DirectMobileChargeResponse)
	_, err := m.session.Post(m.baseURL+"/direct-mobile/charge", &payload, response, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (m *MPower) DirectMobileStatus(token string) (*DirectMobileStatusResponse, error) {
	payload := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	response := new(DirectMobileStatusResponse)
	_, err := m.session.Post(m.baseURL+"/direct-mobile/status", &payload, response, nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
