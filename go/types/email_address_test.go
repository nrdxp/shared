package types

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type emailAddress struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

func TestEmailAddressMarshalJSON(t *testing.T) {
	tests := map[string]struct {
		input emailAddress
		want  string
	}{
		"email marshal - valid": {
			input: emailAddress{
				EmailAddress: "good@example.com",
			},
			want: `{"emailAddress":"good@example.com"}`,
		},
		"email marshal - invalid": {
			input: emailAddress{
				EmailAddress: "",
			},
			want: `{"emailAddress":""}`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tc.input)
			assert.Equal(t, err, nil)
			assert.Equal(t, string(got), tc.want)
		})
	}
}

func TestEmailAddressUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		err   bool
		input string
		want  emailAddress
	}{
		"invalid jsn format": {
			err:   true,
			input: `{"emailAddress":111}`,
			want:  emailAddress{},
		},
		"invalid email format": {
			err:   true,
			input: `{"emailAddress":"badexample.com"}`,
			want:  emailAddress{},
		},
		"good email": {
			input: `{"emailAddress":"Good@example.com"}`,
			want: emailAddress{
				EmailAddress: "good@example.com",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var e emailAddress

			err := json.Unmarshal([]byte(tc.input), &e)
			assert.Equal(t, err != nil, tc.err)
			assert.Equal(t, e, tc.want)
		})
	}
}

func TestEmailAddressValue(t *testing.T) {
	got, err := EmailAddress("Something@example.com").Value()
	assert.Equal(t, err, nil)
	assert.Equal(t, got, "something@example.com")
}

func TestEmailAddressScan(t *testing.T) {
	input := "Something@example.com"

	var e EmailAddress

	err := e.Scan(input)
	assert.Equal(t, err, nil)
	assert.Equal(t, e.String(), strings.ToLower(input))
}
