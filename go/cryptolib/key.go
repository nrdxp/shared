package cryptolib

import (
	"context"
	"crypto"
	"database/sql/driver"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/candiddev/shared/go/cli"
	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
	"github.com/candiddev/shared/go/types"
)

// Algorithm is the type of key.
type Algorithm string

const (
	AlgorithmDefault Algorithm = "default"
)

var ErrParseKeyUnknown = errors.New("unknown key format")
var ErrParseKeyNotImplemented = errors.New("key cannot be used for the required operation")

func newID() string {
	return types.RandString(10)
}

func init() { //nolint:gochecknoinits
	gob.Register(AES128Key(""))
	gob.Register(ECP256PrivateKey(""))
	gob.Register(ECP256PublicKey(""))
	gob.Register(Ed25519PrivateKey(""))
	gob.Register(Ed25519PublicKey(""))
	gob.Register(RSA2048PrivateKey(""))
	gob.Register(RSA2048PublicKey(""))
}

// Key is a generic KeyProvider.
type Key[T KeyProvider] struct {
	ID  string
	Key T
}

// KeyProvider is something that can be a key.
type KeyProvider interface {
	Algorithm() Algorithm
}

// KeyProviderDecryptAsymmetric is a key that decrypts asymmetrically.
type KeyProviderDecryptAsymmetric interface {
	DecryptAsymmetric(input EncryptedValue) (output []byte, err error)
	KeyProvider
}

// KeyProviderEncryptAsymmetric is a key that encrypts asymmetrically.
type KeyProviderEncryptAsymmetric interface {
	EncryptAsymmetric(input []byte) (output EncryptedValue, err error)
	KeyProvider
}

// KeyProviderEncryptSymmetric is a key that encrypts symmetrically.
type KeyProviderEncryptSymmetric interface {
	DecryptSymmetric(input EncryptedValue) (output []byte, err error)
	EncryptSymmetric(input []byte) (output EncryptedValue, err error)
	KeyProvider
}

// KeyProviderSign is a key that signs things.
type KeyProviderSign interface {
	Sign(message []byte, hash crypto.Hash) (signature []byte, err error)
	KeyProvider
}

// KeyProviderVerify is a key that verifies things.
type KeyProviderVerify interface {
	Verify(message []byte, hash crypto.Hash, signature []byte) (err error)
	KeyProvider
}

// ParseKey turns a string into a Key or error.
func ParseKey[T KeyProvider](s string) (Key[T], error) {
	var k Key[T]

	var kp KeyProvider

	if r := strings.Split(s, ":"); len(r) == 2 || len(r) == 3 {
		switch Algorithm(r[0]) {
		case AlgorithmAES128:
			kp = AES128Key(r[1])
		case AlgorithmECP256Private:
			kp = ECP256PrivateKey(r[1])
		case AlgorithmECP256Public:
			kp = ECP256PublicKey(r[1])
		case AlgorithmEd25519Private:
			kp = Ed25519PrivateKey(r[1])
		case AlgorithmEd25519Public:
			kp = Ed25519PublicKey(r[1])
		case AlgorithmRSA2048Private:
			kp = RSA2048PrivateKey(r[1])
		case AlgorithmRSA2048Public:
			kp = RSA2048PublicKey(r[1])
		case AlgorithmNone:
			kp = None(r[1])
		case AlgorithmDefault:
		}

		a, ok := any(kp).(T)
		if !ok {
			return k, fmt.Errorf("%w: %v", ErrParseKeyNotImplemented, reflect.TypeOf(k))
		}

		k.Key = a

		if len(r) == 3 {
			k.ID = r[2]
		}

		return k, nil
	}

	return k, fmt.Errorf("%w: %s", ErrParseKeyUnknown, s)
}

func (k Key[T]) Algorithm() Algorithm {
	return k.Key.Algorithm()
}

// IsNil returns whether the key is nil.
func (k Key[T]) IsNil() bool {
	return any(k.Key) == nil
}

func (k *Key[T]) Scan(src any) error {
	return k.UnmarshalText([]byte(src.(string)))
}

func (k Key[T]) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

func (k Key[T]) String() string {
	if k.IsNil() {
		return ""
	}

	o := fmt.Sprintf("%s:%v", k.Key.Algorithm(), k.Key)
	if k.ID != "" {
		o += ":" + k.ID
	}

	return o
}

func (k Key[T]) Value() (driver.Value, error) {
	return k.MarshalText()
}

func (k *Key[T]) UnmarshalText(data []byte) error {
	var err error

	if len(data) == 0 {
		return nil
	}

	*k, err = ParseKey[T](string(data))

	return err
}

// GenerateKeys is a helper function for CLI apps to generate crypto keys.
func GenerateKeys[T cli.AppConfig[any]]() cli.Command[T] {
	return cli.Command[T]{
		ArgumentsRequired: []string{
			"encrypt-asymmetric, encrypt-symmetric, sign-verify",
		},
		Name: "generate-keys",
		Run: func(ctx context.Context, args []string, _ T) errs.Err {
			m := map[string]string{}

			switch args[1] {
			case "encrypt-symmetric":
				key, err := NewKeyEncryptSymmetric()
				if err != nil {
					return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
				}

				m["key"] = key.String()
			case "encrypt-asymmetric":
				prv, pub, err := NewKeysEncryptAsymmetric()
				if err != nil {
					return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
				}

				m["privateKey"] = prv.String()
				m["publicKey"] = pub.String()
			case "sign-verify":
				prv, pub, err := NewKeysSign()
				if err != nil {
					return logger.Error(ctx, errs.ErrReceiver.Wrap(err))
				}

				m["privateKey"] = prv.String()
				m["publicKey"] = pub.String()
			}

			logger.Raw(types.JSONToString(m) + "\n")

			return nil
		},
		Usage: "Generate cryptographic keys",
	}
}

// KeyDecryptAsymmetric is Key that decrypts asymmetrically.
type KeyDecryptAsymmetric struct {
	Key[KeyProviderDecryptAsymmetric]
}

func (k *KeyDecryptAsymmetric) DecryptAsymmetric(input EncryptedValue) ([]byte, error) {
	return k.Key.Key.DecryptAsymmetric(input)
}

type KeysDecryptAsymmetric []KeyDecryptAsymmetric

// KeyEncryptAsymmetric is Key that encrypts asymmetrically.
type KeyEncryptAsymmetric struct {
	Key[KeyProviderEncryptAsymmetric]
}

func (k *KeyEncryptAsymmetric) EncryptAsymmetric(value []byte) (EncryptedValue, error) {
	o, err := k.Key.Key.EncryptAsymmetric(value)
	o.KeyID = k.ID

	return o, err
}

func NewKeysEncryptAsymmetric() (KeyDecryptAsymmetric, KeyEncryptAsymmetric, error) {
	prv, pub, err := NewRSA2048()
	id := newID()

	return KeyDecryptAsymmetric{
			Key: Key[KeyProviderDecryptAsymmetric]{
				ID:  id,
				Key: prv,
			},
		},
		KeyEncryptAsymmetric{
			Key: Key[KeyProviderEncryptAsymmetric]{
				ID:  id,
				Key: pub,
			},
		}, err
}

// KeyEncryptSymmetric is Key that encrypts Symmetrically.
type KeyEncryptSymmetric struct {
	Key[KeyProviderEncryptSymmetric]
}

func NewKeyEncryptSymmetric() (KeyEncryptSymmetric, error) {
	k, err := NewAES128Key()

	return KeyEncryptSymmetric{
		Key: Key[KeyProviderEncryptSymmetric]{
			ID:  types.RandString(10),
			Key: k,
		},
	}, err
}

func (k *KeyEncryptSymmetric) DecryptSymmetric(input EncryptedValue) ([]byte, error) {
	return k.Key.Key.DecryptSymmetric(input)
}

func (k *KeyEncryptSymmetric) EncryptSymmetric(value []byte) (EncryptedValue, error) {
	o, err := k.Key.Key.EncryptSymmetric(value)
	o.KeyID = k.ID

	return o, err
}

// KeySign is Key that signs things.
type KeySign struct {
	Key[KeyProviderSign]
}

func (k *KeySign) Sign(message []byte, hash crypto.Hash) (signature []byte, err error) {
	return k.Key.Key.Sign(message, hash)
}

// KeyVerify is Key that verifies things.
type KeyVerify struct {
	Key[KeyProviderVerify]
}

// KeysVerify are Keys that verifies things.
type KeysVerify []KeyVerify

func (k *KeyVerify) Verify(message []byte, hash crypto.Hash, signature []byte) (err error) {
	return k.Key.Key.Verify(message, hash, signature)
}

func NewKeysSign() (KeySign, KeyVerify, error) {
	prv, pub, err := NewEd25519()
	id := newID()

	return KeySign{
			Key: Key[KeyProviderSign]{
				ID:  id,
				Key: prv,
			},
		}, KeyVerify{
			Key: Key[KeyProviderVerify]{
				ID:  id,
				Key: pub,
			},
		}, err
}
