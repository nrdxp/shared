package types

import (
	"encoding/json"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

type yearmonth struct {
	YearMonth YearMonth `json:"yearMonth"`
}

func TestYearMonthFromString(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  YearMonth
	}{
		"bad": {
			err:   ErrYearMonth,
			input: "",
			want:  YearMonth(0),
		},
		"good": {
			input: "202112",
			want:  YearMonth(202112),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := YearMonthFromString(tc.input)
			assert.Equal(t, err, tc.err)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestYearMonthMarshalJSON(t *testing.T) {
	got, err := json.Marshal(yearmonth{
		YearMonth: 201910,
	})
	assert.Equal(t, err, nil)
	assert.Equal(t, string(got), `{"yearMonth":201910}`)
}

func TestYearMonthString(t *testing.T) {
	assert.Equal(t, YearMonth(202112).String(), "202112")
}

func TestYearMonthStringDash(t *testing.T) {
	assert.Equal(t, YearMonth(202112).StringDash(), "2021-12")
}

func TestYearMonthUnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		err   error
		input string
		want  yearmonth
	}{
		"invalid1": {
			err:   ErrYearMonth,
			input: `{"yearMonth":201913}`,
			want:  yearmonth{},
		},
		"invalid2": {
			err:   ErrYearMonth,
			input: `{"yearMonth":"201913"}`,
			want:  yearmonth{},
		},
		"invalid3": {
			err:   ErrYearMonth,
			input: `{"yearMonth":201900}`,
			want:  yearmonth{},
		},
		"invalid4": {
			err:   ErrYearMonth,
			input: `{"yearMonth":1010}`,
			want:  yearmonth{},
		},
		"good": {
			input: `{"yearMonth":201910}`,
			want: yearmonth{
				YearMonth: 201910,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var y yearmonth
			assert.Equal(t, json.Unmarshal([]byte(tc.input), &y), tc.err)
			assert.Equal(t, y, tc.want)
		})
	}
}
