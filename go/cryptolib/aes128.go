package cryptolib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sync"
)

// AES128Key is a key used for AES encryption.
type AES128Key string

const (
	AlgorithmAES128     Algorithm  = "aes128"
	EncryptionAES128GCM Encryption = "aes128gcm"
)

var ErrAES128CiphertextLength = errors.New("length of ciphertext is too short, probably invalid")

var aesKeys = struct { //nolint: gochecknoglobals
	keys  map[AES128Key]cipher.Block
	mutex sync.Mutex
}{
	keys: map[AES128Key]cipher.Block{},
}

// NewAES128Key generates a new AES key or an error.
func NewAES128Key() (AES128Key, error) {
	key := make([]byte, 16) // 16 bytes is an AES128 key

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", fmt.Errorf("%w: %w", ErrGeneratingKey, err)
	}

	return AES128Key(base64.StdEncoding.EncodeToString(key)), nil
}

func (AES128Key) Algorithm() Algorithm {
	return AlgorithmAES128
}

func (k AES128Key) DecryptSymmetric(v EncryptedValue) ([]byte, error) {
	if v.Encryption == EncryptionDefault || v.Encryption == EncryptionAES128GCM {
		b, err := base64.StdEncoding.DecodeString(v.Ciphertext)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrDecodingValue, err)
		}

		return k.DecryptGCM(b)
	}

	return nil, v.ErrUnsupportedDecrypt()
}

func (k AES128Key) DecryptGCM(value []byte) ([]byte, error) {
	b, err := k.Key()
	if err != nil {
		return nil, err
	}

	if len(value) < 13 {
		return nil, ErrAES128CiphertextLength
	}

	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGeneratingGCM, err)
	}

	output, err := gcm.Open(nil, value[:12], value[12:], nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecryptingKey, err)
	}

	return output, nil
}

func (k AES128Key) EncryptSymmetric(value []byte) (EncryptedValue, error) {
	v, err := k.EncryptGCM(value)

	return EncryptedValue{
		Ciphertext: base64.StdEncoding.EncodeToString(v),
		Encryption: EncryptionAES128GCM,
	}, err
}

func (k AES128Key) EncryptGCM(value []byte) ([]byte, error) {
	b, err := k.Key()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGeneratingNonce, err)
	}

	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGeneratingGCM, err)
	}

	output := gcm.Seal(nonce, nonce, value, nil)

	return output, nil
}

func (k AES128Key) Key() (cipher.Block, error) {
	aesKeys.mutex.Lock()

	defer aesKeys.mutex.Unlock()

	var ok bool

	var b cipher.Block

	if b, ok = aesKeys.keys[k]; !ok {
		bytesKey, err := base64.StdEncoding.DecodeString(string(k))
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrDecodingKey, err)
		}

		b, err = aes.NewCipher(bytesKey)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrCreatingCipher, err)
		}

		aesKeys.keys[k] = b
	}

	return b, nil
}
