package assert

import (
	"errors"
	"testing"
)

type mockJSON struct {
	A string `json:"a"`
	B bool   `json:"b"`
	c string
}

type mockStruct struct {
	A string
	B bool
}

type mockTest struct{}

var fail bool

var m = &mockTest{}

func (*mockTest) Errorf(_ string, _ ...any) {
	fail = true
}

func (*mockTest) Helper() {
	fail = false
}

func TestContains(t *testing.T) {
	tests := map[string]struct {
		fail bool
		got  string
		want string
	}{
		"contains": {
			got:  "a string",
			want: "string",
		},
		"not contains": {
			fail: true,
			got:  "a string",
			want: "A",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			Contains(m, tc.got, tc.want)

			if fail != tc.fail {
				t.Errorf("got %t, want %t", fail, tc.fail)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := map[string]struct {
		fail bool
		got  mockStruct
		want mockStruct
	}{
		"equal": {
			got: mockStruct{
				A: "A",
				B: true,
			},
			want: mockStruct{
				A: "A",
				B: true,
			},
		},
		"not equal": {
			fail: true,
			got: mockStruct{
				A: "A",
				B: false,
			},
			want: mockStruct{
				A: "A",
				B: true,
			},
		},
		"got nil": {
			fail: true,
			want: mockStruct{
				A: "A",
				B: true,
			},
		},
		"want nil": {
			fail: true,
			got: mockStruct{
				A: "A",
				B: true,
			},
		},
		"both nil": {},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			Equal(m, tc.got, tc.want)

			if fail != tc.fail {
				t.Errorf("got %t, want %t", fail, tc.fail)
			}
		})
	}
}

func TestEqualJSON(t *testing.T) {
	tests := map[string]struct {
		fail bool
		got  any
		want any
	}{
		"equal": {
			got: mockJSON{
				A: "A",
				B: true,
				c: "c",
			},
			want: mockJSON{
				A: "A",
				B: true,
				c: "d",
			},
		},
		"not equal": {
			fail: true,
			got: mockJSON{
				A: "A",
				B: false,
				c: "c",
			},
			want: mockJSON{
				A: "A",
				B: true,
				c: "c",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			EqualJSON(m, tc.got, tc.want)

			if fail != tc.fail {
				t.Errorf("got %t, want %t", fail, tc.fail)
			}
		})
	}
}

func TestHasErr(t *testing.T) {
	err := errors.New("test")

	tests := map[string]struct {
		fail bool
		got  error
		want error
	}{
		"nil,nil": {},
		"err,nil": {
			got:  err,
			fail: true,
		},
		"err,want": {
			got:  err,
			want: err,
		},
		"nil,want": {
			want: err,
			fail: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			HasErr(m, tc.got, tc.want)

			if fail != tc.fail {
				t.Errorf("got %t, want %t", fail, tc.fail)
			}
		})
	}
}
