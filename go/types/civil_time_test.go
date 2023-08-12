package types

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/candiddev/shared/go/assert"
)

var ct = CivilTime{
	Hour:   15,
	Minute: 4,
}

type ctS struct {
	CivilTime CivilTime `json:"civilTime"`
}

func TestCivilTimeOf(t *testing.T) {
	cs := fmt.Sprintf("%02d:%02d", ct.Hour, ct.Minute)
	cn, _ := time.Parse("15:04", cs)
	assert.Equal(t, CivilTimeOf(cn), CivilTime{
		Hour:   15,
		Minute: 4,
	})
}

func TestParseCivilTime(t *testing.T) {
	tests := map[string]struct {
		err   bool
		input string
		want  CivilTime
	}{
		"bad - too short": {
			err:   true,
			input: "15",
			want:  CivilTime{},
		},
		"bad - too long": {
			err:   true,
			input: "00:111",
			want:  CivilTime{},
		},
		"bad - hour": {
			err:   true,
			input: "24:00",
			want:  CivilTime{},
		},
		"bad - minute": {
			err:   true,
			input: "00:60",
			want:  CivilTime{},
		},
		"good": {
			input: "15:04",
			want: CivilTime{
				Hour:   15,
				Minute: 4,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCivilTime(tc.input)
			assert.Equal(t, err != nil, tc.err)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestCivilTimeString(t *testing.T) {
	tests := map[string]struct {
		input CivilTime
		want  string
	}{
		"am": {
			input: CivilTime{
				Hour:   7,
				Minute: 7,
			},
			want: "7:07 AM",
		},
		"midnight": {
			input: CivilTime{
				Hour:   0,
				Minute: 7,
			},
			want: "12:07 AM",
		},
		"noon": {
			input: CivilTime{
				Hour:   12,
				Minute: 7,
			},
			want: "12:07 PM",
		},
		"pm": {
			input: CivilTime{
				Hour:   15,
				Minute: 45,
			},
			want: "3:45 PM",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.input.String12(), tc.want)
		})
	}
}

func TestCivilTimeString12(t *testing.T) {
	assert.Equal(t, ct.String12(), "3:04 PM")
}

func TestCivilTimeAddMinutes(t *testing.T) {
	assert.Equal(t, ct.AddMinutes(30).Minute, ct.Minute+30)
}

func TestCivilTimeMarshalJSON(t *testing.T) {
	var m []byte

	m, err := json.Marshal(ctS{
		CivilTime: CivilTime{
			Hour:   12,
			Minute: 0,
		},
	},
	)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(m), `{"civilTime":"12:00"}`)

	m, err = json.Marshal(ctS{
		CivilTime: CivilTime{},
	})
	assert.Equal(t, err, nil)
	assert.Equal(t, string(m), `{"civilTime":"00:00"}`)
}

func TestCivilTimeUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		err   bool
		input string
		want  ctS
	}{
		"bad time": {
			err:   true,
			input: `{"civilTime":"55:90"}`,
			want:  ctS{},
		},
		"null time": {
			input: `{"civilTime":null}`,
			want:  ctS{},
		},
		"good time": {
			input: `{"civilTime":"12:00"}`,
			want: ctS{
				CivilTime: CivilTime{
					Hour:   12,
					Minute: 0,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var cn ctS

			err := json.Unmarshal([]byte(tc.input), &cn)
			assert.Equal(t, err != nil, tc.err)
			assert.Equal(t, cn, tc.want)
		})
	}
}

func TestCivilTimeValue(t *testing.T) {
	got, err := ct.Value()
	assert.Equal(t, err, nil)
	assert.Equal(t, got, "15:04")

	got, err = CivilTime{}.Value()
	assert.Equal(t, err, nil)
	assert.Equal(t, got, "00:00")
}

func TestCivilTimeScan(t *testing.T) {
	now := time.Now()

	var cn CivilTime

	assert.Equal(t, cn.Scan(now), nil)
	assert.Equal(t, cn, CivilTimeOf(now))
}
