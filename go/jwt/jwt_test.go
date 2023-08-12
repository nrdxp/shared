package jwt

import (
	"testing"
	"time"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/crypto"
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

func TestSignJWT(t *testing.T) {
	prvStr, pubStr, _ := crypto.NewEd25519()
	d := time.Now().Add(-24 * time.Hour)
	expires := time.Now().Add(24 * time.Hour)

	c := jwtCustom{
		Licensed:     true,
		Name:         "a",
		LastPurchase: d,
	}

	tokenBad, err := SignJWT(prvStr, &c, d, "https://example.com", "https://example.com", "Account")
	assert.Equal(t, err, nil)

	tokenGood, err := SignJWT(prvStr, &c, expires, "https://example.com", "https://example.com", "Account")
	assert.Equal(t, err, nil)

	var cOutBad jwtCustom

	var cOutGood jwtCustom

	_, err = VerifyJWT(pubStr, &cOutBad, tokenBad)

	assert.Contains(t, err.Error(), "error parsing token: token is expired")

	e, err := VerifyJWT(pubStr, &cOutGood, tokenGood)

	assert.Equal(t, err, nil)
	assert.Equal(t, e, expires.Truncate(time.Second))

	assert.Equal(t, cOutGood, c)
}
