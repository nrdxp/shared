package cryptolib

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"github.com/candiddev/shared/go/types"
)

// Encryption is the type of encryption.
type Encryption string

const EncryptionDefault = "default"

// EncryptedValue is a decoded encrypted value.
type EncryptedValue struct {
	Ciphertext string
	Encryption Encryption
	KeyID      string
}

// ParseEncryptedValue turns a string into an EncryptedValue or error.
func ParseEncryptedValue(s string) (EncryptedValue, error) {
	if r := strings.Split(s, ":"); len(r) == 2 || len(r) == 3 {
		var e Encryption

		switch Encryption(r[0]) {
		case EncryptionNone:
			e = EncryptionNone
		case EncryptionAES128GCM:
			e = EncryptionAES128GCM
		case EncryptionRSA2048OAEPSHA256:
			e = EncryptionRSA2048OAEPSHA256
		}

		if e != "" {
			v := EncryptedValue{
				Ciphertext: r[1],
				Encryption: e,
			}

			if len(r) == 3 {
				v.KeyID = r[2]
			}

			return v, nil
		}
	}

	return EncryptedValue{}, ErrUnknownEncryption
}

func (e EncryptedValue) ErrUnsupportedDecrypt() error {
	return fmt.Errorf("%s: %w", e.Encryption, ErrUnsupportedDecrypt)
}

func (e EncryptedValue) MarshalJSON() ([]byte, error) {
	output := ""

	if e != (EncryptedValue{}) {
		output = strconv.Quote(e.String())
	}

	return []byte(output), nil
}

func (e *EncryptedValue) String() string {
	o := fmt.Sprintf("%s:%s", e.Encryption, e.Ciphertext)
	if e.KeyID != "" {
		o += ":" + e.KeyID
	}

	return o
}

func (e *EncryptedValue) UnmarshalJSON(data []byte) error {
	var err error

	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	*e, err = ParseEncryptedValue(s)

	return err
}

func (e EncryptedValue) Value() (driver.Value, error) {
	if e.Encryption == "" {
		return "", nil
	}

	return e.String(), nil
}

func (e *EncryptedValue) Scan(src any) error {
	var err error

	if src != nil {
		*e, err = ParseEncryptedValue(src.(string))
	}

	return err
}

// EncryptedValues is multiple EncryptedValue.
type EncryptedValues []EncryptedValue

func (e EncryptedValues) SliceString() types.SliceString {
	s := types.SliceString{}

	for i := range e {
		s = append(s, e[i].String())
	}

	return s
}

func (e EncryptedValues) MarshalJSON() ([]byte, error) {
	return e.SliceString().MarshalJSON()
}

func (e EncryptedValues) Value() (driver.Value, error) {
	return e.SliceString().Value()
}

func (e *EncryptedValues) Scan(src any) error {
	if src != nil {
		source := string(src.([]byte))
		if source != "" && source != "{}" && source != "{NULL}" {
			output := strings.TrimRight(strings.TrimLeft(source, "{"), "}")
			array := strings.Split(output, ",")

			slice := EncryptedValues{}

			for i := range array {
				var s string

				if array[i] != `""` && array[i][0] == '"' {
					s = array[i][1 : len(array[i])-1]
				} else {
					s = array[i]
				}

				v, err := ParseEncryptedValue(s)
				if err != nil {
					return err
				}

				slice = append(slice, v)
			}

			*e = slice

			return nil
		}
	}

	*e = EncryptedValues{}

	return nil
}
