package crypto

import (
	"fmt"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestParseEncryptedValue(t *testing.T) {
	tests := map[string]struct {
		err   error
		input Type
		want  string
	}{
		"None": {
			input: TypeNone + "$none",
			want:  "none$none",
		},
		"AES128": {
			input: TypeAES128GCM + "$aes",
			want:  "aes128gcm$aes",
		},
		"RSA2048": {
			input: TypeRSA2048 + "$rsa",
			want:  "rsa2048$rsa",
		},
		"Unknown": {
			input: Type("unknown$unknown"),
			err:   ErrUnknownEncryption,
			want:  "$",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseEncryptedValue(string(tc.input))
			assert.Equal(t, err, tc.err)
			assert.Equal(t, got.String(), tc.want)
		})
	}
}

func TestEncryptedValue(t *testing.T) {
	aesKey, _ := NewAESKey()
	prv, pub, _ := NewRSA()
	value := "testing"

	tests := map[string]struct {
		err             error
		inputEncryptKey Encrypter
		inputDecryptKey Decrypter
		inputValue      string
	}{
		"none": {
			inputValue: value,
		},
		"aes128gcm": {
			inputDecryptKey: aesKey,
			inputEncryptKey: aesKey,
			inputValue:      value,
		},
		"rsa2048": {
			inputDecryptKey: prv,
			inputEncryptKey: pub,
			inputValue:      value,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ev, err := EncryptValue(tc.inputEncryptKey, tc.inputValue)
			assert.Equal(t, err, tc.err)

			if err != nil {
				return
			}

			got, err := ev.Decrypt(tc.inputDecryptKey)
			assert.Equal(t, err, nil)
			assert.Equal(t, got, value)

			// Marshal
			strout := fmt.Sprintf("%s$%s", ev.Encryption, ev.Ciphertext)
			bytout := []byte(fmt.Sprintf(`"%s"`, strout))
			jsonout, _ := ev.MarshalJSON()
			assert.Equal(t, bytout, jsonout)

			// Unmarshal
			var evout EncryptedValue

			evout.UnmarshalJSON(bytout)
			assert.Equal(t, evout, ev)

			// Value
			drvout, _ := ev.Value()
			assert.Equal(t, drvout.(string), strout)

			// Scan
			evout = EncryptedValue{}
			evout.Scan(strout)
			assert.Equal(t, evout, ev)
		})
	}
}

func TestEncryptedValues(t *testing.T) {
	ev := EncryptedValues{
		{
			Ciphertext: "test1",
			Encryption: TypeNone,
		},
		{
			Ciphertext: "test2",
			Encryption: TypeNone,
		},
	}

	bytout := []byte(`["none$test1","none$test2"]`)
	jsonout, _ := ev.MarshalJSON()
	assert.Equal(t, jsonout, bytout)

	vout := []byte("{none$test1,none$test2}")
	valout, _ := ev.Value()
	assert.Equal(t, valout.([]byte), vout)

	var evout EncryptedValues

	evout.Scan(vout)

	assert.Equal(t, evout, ev)
}
