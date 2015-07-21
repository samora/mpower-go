package mpower

// Setup has credentials of current merchant.
type Setup struct {
	MasterKey  string `valid:"required"`
	PrivateKey string `valid:"required"`
	PublicKey  string `valid:"required"`
	IsLive     bool
}
