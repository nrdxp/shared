// Package cryptolib is a crypto wrapper that provides future-proof crypto abstractions.
package cryptolib

import (
	"errors"
)

var (
	ErrCreatingCipher       = errors.New("error creating cipher")
	ErrCreatingHash         = errors.New("error creating hash")
	ErrDecodingKey          = errors.New("error decoding key")
	ErrDecodingPrivateKey   = errors.New("error decoding base64 PKCS8 private key")
	ErrDecodingPublicKey    = errors.New("error decoding base64 PKIX public key")
	ErrDecodingValue        = errors.New("error decoding base64 value")
	ErrDecryptingKey        = errors.New("error decrypting with key")
	ErrDecryptingPrivateKey = errors.New("error decrypting with private key")
	ErrEncryptingPublicKey  = errors.New("error encrypting with public key")
	ErrGeneratingGCM        = errors.New("error generating GCM")
	ErrGeneratingKey        = errors.New("error generating key")
	ErrGeneratingNonce      = errors.New("error generating nonce")
	ErrGeneratingPrivateKey = errors.New("error generating private key")
	ErrMarshalingPublicKey  = errors.New("error marshaling public key")
	ErrMarshalingPrivateKey = errors.New("error marshaling private key")
	ErrNoPrivateKey         = errors.New("no private key found")
	ErrNoPublicKey          = errors.New("no public key found")
	ErrParsingPrivateKey    = errors.New("error parsing private key")
	ErrParsingPublicKey     = errors.New("error parsing public key")
	ErrSign                 = errors.New("error signing message")
	ErrSignKeyFormat        = errors.New("key does not match format <type>:<base64PKCS8PrivateKey>")
	ErrVerify               = errors.New("error verifying signature against message")
	ErrUnknownEncryption    = errors.New("unknown encryption")
	ErrUnsupportedDecrypt   = errors.New("decrypter doesn't support encryption type")
)
