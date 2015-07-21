package mpower

// Store has metadata of current merchant.
type Store struct {
	Name          string `valid:"required"`
	Tagline       string
	PhoneNumber   string
	PostalAddress string
	LogoURL       string
}
