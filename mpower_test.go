package mpower

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMPower(t *testing.T) {
	setup := Setup{
		MasterKey:  "",
		PrivateKey: "",
		PublicKey:  "",
		Token:      "",
		IsLive:     true,
	}
	store := Store{
		Name: "Demo Store",
	}
	m := New(nil, setup, store)
	directPayResponse, err := m.DirectPay(DirectPayPayload{
		AccountAlias: "samora",
		Amount:       10,
	})
	assert.NoError(t, err)
	assert.True(t, directPayResponse.IsSuccess())

	directMobileCharge, err := m.DirectMobileCharge(DirectMobileChargePayload{
		CustomerName:   "Samora",
		CustomerEmail:  "samora@example.com",
		CustomerPhone:  "0561516300",
		WalletProvider: "airtel",
		Amount:         10,
	})
	assert.NoError(t, err)
	assert.True(t, directMobileCharge.IsSuccess())
}
