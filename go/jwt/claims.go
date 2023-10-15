package jwt

import (
	"encoding/json"
	"strings"
)

// RegisteredClaims are default fields in a JWT.
type RegisteredClaims struct {
	Audience  Audience `json:"aud,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	ID        string   `json:"jti,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	Subject   string   `json:"sub,omitempty"`
}

// CustomClaims is a struct that can be parsed from a JWT payload.
type CustomClaims interface {
	GetRegisteredClaims() *RegisteredClaims
	Valid() error
}

// Audience is a string or array of strings.
type Audience []string

// MarshalText converts an audience string to JSON array or singleton.
func (a Audience) MarshalText() ([]byte, error) {
	switch {
	case len(a) == 0:
		return nil, nil
	case len(a) == 1:
		return []byte(a[0]), nil
	}

	return json.Marshal([]string(a))
}

// UnmarshalText converts an audience string to JSON array or singleton.
func (a *Audience) UnmarshalText(data []byte) error {
	if data == nil {
		return nil
	}

	if strings.HasPrefix(string(data), "[") {
		as := []string{}
		if err := json.Unmarshal(data, &as); err != nil {
			return err
		}

		*a = Audience(as)
	} else {
		*a = Audience{string(data)}
	}

	return nil
}
