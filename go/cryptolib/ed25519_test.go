package cryptolib

import (
	"crypto/ed25519"
	"encoding/base64"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestEd25519(t *testing.T) {
	prvStr, pubStr, err := NewEd25519()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(prvStr), 64)
	assert.Equal(t, len(pubStr), 60)

	badPrv1 := Ed25519PrivateKey("asd")
	_, err = badPrv1.PrivateKey()
	assert.HasErr(t, err, ErrDecodingPrivateKey)

	badPrv2 := Ed25519PrivateKey(pubStr)
	_, err = badPrv2.PrivateKey()
	assert.HasErr(t, err, ErrParsingPrivateKey)

	prvKey, err := prvStr.PrivateKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, ed25519PrivateKeys.mutex.TryLock(), true)
	ed25519PrivateKeys.mutex.Unlock()
	assert.Equal(t, ed25519PrivateKeys.keys[prvStr], prvKey)
	assert.Equal(t, len(ed25519PrivateKeys.keys), 1)
	ed25519PrivateKeys.keys = map[Ed25519PrivateKey]ed25519.PrivateKey{}

	badPub1 := Ed25519PublicKey("asd")
	_, err = badPub1.PublicKey()
	assert.HasErr(t, err, ErrDecodingPublicKey)

	badPub2 := Ed25519PublicKey(prvStr)
	_, err = badPub2.PublicKey()
	assert.HasErr(t, err, ErrParsingPublicKey)

	pubKey, err := pubStr.PublicKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, ed25519PublicKeys.mutex.TryLock(), true)
	ed25519PublicKeys.mutex.Unlock()
	assert.Equal(t, len(ed25519PublicKeys.keys), 1)
	assert.Equal(t, ed25519PublicKeys.keys[pubStr], pubKey)
	ed25519PublicKeys.keys = map[Ed25519PublicKey]ed25519.PublicKey{}

	// Sign/Verify
	m := []byte("hello")
	sig, err := prvStr.Sign(m, 0)
	assert.HasErr(t, err, nil)
	err = pubStr.Verify(m, 0, sig)
	assert.HasErr(t, err, nil)

	prvStr = Ed25519PrivateKey("OkyNtJfJqnDrzyi7UwOo6OG7A2GuA4UFUWowvGADtGBbfp5oeDulsl4TZ5uSTs9U9MQDbg9K+B4OaDY1oTgVTQ==")
	pubStr = Ed25519PublicKey("W36eaHg7pbJeE2ebkk7PVPTEA24PSvgeDmg2NaE4FU0=")
	sig, err = prvStr.Sign(m, 0)
	assert.HasErr(t, err, nil)
	err = pubStr.Verify(m, 0, sig)
	assert.HasErr(t, err, nil)

	b, _ := base64.StdEncoding.DecodeString("iUfnOZ+Wx/j6eS+BVXrlPQDzkUsy0/diZzY8OygbE7vyunKDFhx3P13GHAhA+HsIKDyHNrY3N9yTxH/ieOTJDw==")
	err = pubStr.Verify(m, 0, b)
	assert.HasErr(t, err, nil)
}
