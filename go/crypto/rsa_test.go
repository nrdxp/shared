package crypto

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestRSA(t *testing.T) {
	// Generate keys
	private, public, err := NewRSA()
	assert.Equal(t, err, nil)

	// Test encryption
	got, err := EncryptValue(public, "testing")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(got.Ciphertext), 344)

	gotString, err := got.Decrypt(private)
	assert.Equal(t, err, nil)
	assert.Equal(t, gotString, "testing")

	// Test JavaScript outputs
	private = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCzXQkWaXZEkiRud1HrrgSSYUqwfYGlsBu1weHakj2DrMJhmhVaGSVM98zeJVTBHn476KSFQKqFly7/46Rjzv09k15A6o8EBKlIKiUctWWRV5YJxFoE0L5+iv2dxc53dsP2n4c31htZ2NdtPL1NL7w0GIjhN1/q1drVedt5IK/4W+pnG8VCq/E2r+LR2whsAoc8jWE+UElu9gazGDbEQlBRN4noKCdwsa+Mqk4JJP0myJkK3R8qUewAEXsrliuRQwKEJzKirviluUiPAtgSMWZ54Mj0LzkNgWJ7t1gVDlaZBjZAXC518e03m5aZn3rrfUKqY8kqJGO4WpsqjeMjKEwXAgMBAAECggEACq/zd6tnCaTvmMlxTyq6SQKVEbkYvYelzwmzbHOkSLFFj80/knPHHjPxuBbfWvZ3LRiqnOoCfx1IBqVIyU2vFzktt46CYQnqEEu6lGBWofdnEEPr3TeW8jFT9U+xArtH5+0PgIs4a/0YsQCagvviKfSAS3MiOvrtJCUzp58D/TdbCZffX++It08/tdHhqclA0hAlwD/FtIU8ZfM6iqFKz9j87DAcXW+LRTfeLFOkuqW8GZKuQrT/dj052cuG7PZTJkCyT03OkIKKRr2CSBFP/7WOjB5OuqT6lBtJtMEwhjTVsMHRdwiEc8owT+3Ujl6Ud+O7eWUJmcy8NB13EXDPeQKBgQC4sTuXNNiz8DBJrNC7hvjhR2zFPoJ1DNPKbKg9dvjQyi33+NOR1CvPyOzYvkwwVthNwyEoHGxb33OZjykVdUZf5u/z0LJfw8TWaGmWy4YeMTRE/k10fKdu7aVRZe1bTwMmKIUqbmrodkZzmP5N2xiZGSpNXuLdHqggaOji2rqEowKBgQD4nRnfRF/hxKHiGufZ7fV7hXxmvbAza/icsn/wAIfhIfYuJUgXpACdsuIDaDp57d4mvqtCScyCWEuJLF1OMat35g9IyO6kZl+Ds6Az3Vh6wua2l31Ha7MghX2/nhdCyStBjMKc6wLjscDn3MK1mysdLdO3HV2E10x4/wcP9Qxd/QKBgG4QsQKba8lQCnbdlkcrWIZomlmxtu0qh8zKlDiM/hrvYhW1I+B0IOfjRcKlJaJpHt64urm8tpbJw4pv1nwDN71PUxBp/uMzN8rOzp4fC/5SNGQpaN9SqjVhSnQH2jvgCnBjFXkD9JI0aifZtTOQrkgOxzb0pR5BdtBZtO/8icshAoGANXoNg6IAgd6695jUAGOYJz66WJs62ITEEGiw7l4e3EVi3+AzFqq3WzSMOzFjSrDo1Sq0Sej37Z2yecoqM9SpTnwVE5HtCz3eE8VCk4VsL9PWrrX9Q6yd9kPgnaQS2tHM1SfDAN/oWjP1C6rH0yqSZo9h+ASABkANegVZboxGYb0CgYBpFpRUPMPSalRF07B0fUTijYm5NiSiLU5AthhOeqi6XbV637MoQgE49A3Tv+j8uzXA91vaQ7jW8LM/3ld1e+QEpFTyMTl0LQ8zestfhuyrOnjkes7Xz1pOwlaSOCJ/p++hSV2tv8uz+tbyVzvnnpUUCADvUktMr/m3HvqVc//MuQ=="
	public = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs10JFml2RJIkbndR664EkmFKsH2BpbAbtcHh2pI9g6zCYZoVWhklTPfM3iVUwR5+O+ikhUCqhZcu/+OkY879PZNeQOqPBASpSColHLVlkVeWCcRaBNC+for9ncXOd3bD9p+HN9YbWdjXbTy9TS+8NBiI4Tdf6tXa1XnbeSCv+FvqZxvFQqvxNq/i0dsIbAKHPI1hPlBJbvYGsxg2xEJQUTeJ6CgncLGvjKpOCST9JsiZCt0fKlHsABF7K5YrkUMChCcyoq74pblIjwLYEjFmeeDI9C85DYFie7dYFQ5WmQY2QFwudfHtN5uWmZ96631CqmPJKiRjuFqbKo3jIyhMFwIDAQAB"
	jsOutput, _ := ParseEncryptedValue(string(TypeRSA2048) + "$WIuY/t9cIWH/BrUf6FzEGpTD8n5wCYix4Ykf0vNZ2QlQrN9CWACYAddjKPMKFudHSqh12DrWS65M85NdA+g9wMaWtNyJ60SYvyCLmjqdW1sLCHFnRJO+mvBbZZhysbvx3//bhd9MbrKJ/9kE3PCdJ++bEx/QP/npXIErv1mvfV1h7f88uUm2nuT7j3e80hEMa9BY5ajbAJ0yJCUyaHhHn331Q1MhtboehzD1/hHaE/ja6ve6diOR9McsuRMmrj68jmoZcX+lscivGUym1uD8E0/ox2AX92V+mTDZHqjDdbl7KrCTQvOXRb/1Ib2viLhg6hK7kftjBEEUOxV+czGHmQ==")

	got, err = EncryptValue(public, "testing")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(got.Ciphertext), 344)

	gotString, err = got.Decrypt(private)
	assert.Equal(t, err, nil)
	assert.Equal(t, gotString, "testing")

	gotString, err = jsOutput.Decrypt(private)
	assert.Equal(t, err, nil)
	assert.Equal(t, gotString, "testing")
}
