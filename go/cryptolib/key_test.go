package cryptolib

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestNewKeyEncryptSymmetric(t *testing.T) {
	k, err := NewKeyEncryptSymmetric()
	assert.HasErr(t, err, nil)
	assert.Equal(t, k.Key.Algorithm(), AlgorithmAES128)
}

func TestNewKeyEncryptAsymmetric(t *testing.T) {
	prv, pub, err := NewKeysEncryptAsymmetric()
	assert.HasErr(t, err, nil)
	assert.Equal(t, prv.Key.Algorithm(), AlgorithmRSA2048Private)
	assert.Equal(t, pub.Key.Algorithm(), AlgorithmRSA2048Public)
}

func TestNewKeysSign(t *testing.T) {
	prv, pub, err := NewKeysSign()
	assert.HasErr(t, err, nil)
	assert.Equal(t, prv.Key.Algorithm(), AlgorithmEd25519Private)
	assert.Equal(t, pub.Key.Algorithm(), AlgorithmEd25519Public)
}

func TestParseKey(t *testing.T) {
	aes, _ := NewKeyEncryptSymmetric()

	k, err := ParseKey[KeyProviderEncryptSymmetric](aes.String())
	assert.HasErr(t, err, nil)
	assert.Equal(t, k.ID, aes.ID)
	assert.Equal(t, k.String(), aes.String())

	_, err = ParseKey[KeyEncryptAsymmetric](aes.String())
	assert.HasErr(t, err, ErrParseKeyNotImplemented)

	_, err = ParseKey[KeyEncryptAsymmetric]("hello!")
	assert.HasErr(t, err, ErrParseKeyUnknown)
}

func TestKey(t *testing.T) {
	aes, _ := NewKeyEncryptSymmetric()

	j, err := aes.MarshalText()
	assert.HasErr(t, err, nil)
	assert.Equal(t, string(j), aes.String())

	assert.Equal(t, aes.String(), fmt.Sprintf("%s:%s:%s", aes.Key.Algorithm(), aes.Key.Key, aes.ID))

	var s KeyDecryptAsymmetric

	assert.Equal(t, s.String(), "")

	k := KeyEncryptSymmetric{}
	k.UnmarshalText(j)
	assert.Equal(t, k, aes)
}

func TestKeys(t *testing.T) {
	_, v, _ := NewKeysSign()

	k := KeysVerify{
		v,
	}

	j, err := json.Marshal(k)

	assert.HasErr(t, err, nil)
	assert.Equal(t, string(j), fmt.Sprintf(`["%s"]`, v.String()))

	n := KeysVerify{}

	assert.HasErr(t, json.Unmarshal(j, &n), nil)
	assert.Equal(t, n, k)
}
