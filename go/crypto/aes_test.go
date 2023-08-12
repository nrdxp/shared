package crypto

import (
	"math"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestAES(t *testing.T) {
	key, err := NewAESKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, int(math.Ceil(float64(16)/3)*4), len(key))

	input := "testing"

	output, err := EncryptValue(key, input)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(math.Ceil((float64(len(input))+12+4+12)/3)*4), len(output.Ciphertext)) // IV + counter + append IV
	assert.Equal(t, TypeAES128GCM, output.Encryption)

	out, err := output.Decrypt(key)
	assert.Equal(t, err, nil)
	assert.Equal(t, out, input)

	// Test JavaScript outputs
	key = "lQQyBeRmGYECoYkafl+4VQ=="
	jsOutput, _ := ParseEncryptedValue(string(TypeAES128GCM) + "$Guqhksdu4nsd6FIrONsXnXyDl8wS78amak0uMZ49KJSly5w=")

	output, err = EncryptValue(key, input)
	assert.Equal(t, err, nil)
	assert.Equal(t, int(math.Ceil((float64(len(input))+12+4+12)/3)*4), len(output.Ciphertext)) // IV + counter + append IV

	out, err = output.Decrypt(key)
	assert.Equal(t, err, nil)
	assert.Equal(t, out, input)

	out, err = jsOutput.Decrypt(key)
	assert.Equal(t, err, nil)
	assert.Equal(t, out, input)
}
