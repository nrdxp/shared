package jwt

import (
	"time"
)

type jwtCustom struct {
	Licensed     bool
	Name         string `json:"name"`
	LastPurchase time.Time
	RegisteredClaims
}

func (j *jwtCustom) GetRegisteredClaims() *RegisteredClaims {
	return &j.RegisteredClaims
}

func (*jwtCustom) Valid() error {
	return nil
}
