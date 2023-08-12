package crypto

import (
	"crypto/ed25519"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestEd25519(t *testing.T) {
	prvStr, pubStr, err := NewEd25519()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(prvStr), 64)
	assert.Equal(t, len(pubStr), 60)

	badPrv1 := Ed25519PrivateKey("asd")
	_, err = badPrv1.decode()
	assert.HasErr(t, err, ErrEd25519DecodingPrivateKey)

	badPrv2 := Ed25519PrivateKey(pubStr)
	_, err = badPrv2.decode()
	assert.HasErr(t, err, ErrEd25519ParsingPrivateKey)

	prvKey, err := prvStr.decode()
	assert.Equal(t, err, nil)
	assert.Equal(t, ed25519PrivateKeys.mutex.TryLock(), true)
	ed25519PrivateKeys.mutex.Unlock()
	assert.Equal(t, ed25519PrivateKeys.keys[&prvStr], prvKey)
	assert.Equal(t, len(ed25519PrivateKeys.keys), 1)
	ed25519PrivateKeys.keys = map[*Ed25519PrivateKey]ed25519.PrivateKey{}

	badPub1 := Ed25519PublicKey("asd")
	_, err = badPub1.decode()
	assert.HasErr(t, err, ErrEd25519DecodingPublicKey)

	badPub2 := Ed25519PublicKey(prvStr)
	_, err = badPub2.decode()
	assert.HasErr(t, err, ErrEd25519ParsingPublicKey)

	pubKey, err := pubStr.decode()
	assert.Equal(t, err, nil)
	assert.Equal(t, ed25519PublicKeys.mutex.TryLock(), true)
	ed25519PublicKeys.mutex.Unlock()
	assert.Equal(t, len(ed25519PublicKeys.keys), 1)
	assert.Equal(t, ed25519PublicKeys.keys[&pubStr], pubKey)
	ed25519PublicKeys.keys = map[*Ed25519PublicKey]ed25519.PublicKey{}
}
