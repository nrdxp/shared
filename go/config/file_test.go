package config

import (
	"context"
	"os"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestGetFile(t *testing.T) {
	logger.UseTestLogger(t)

	stdin := os.Stdin

	tests := map[string]struct {
		err       bool
		extension string
		input     string
		stdin     string
		want      bool
	}{
		"missing": {
			err:   true,
			input: "testdata/missing.json",
			want:  true,
		},
		"invalid json": {
			err:   true,
			input: "testdata/invalid.json",
			want:  true,
		},
		"invalid yaml": {
			err:   true,
			input: "testdata/invalid.yaml",
			want:  true,
		},
		"good json": {
			input: "testdata/good.json",
			want:  false,
		},
		"good yaml": {
			extension: "yaml",
			input:     "testdata/good.yaml",
			want:      true,
		},
		"wrong extension": {
			extension: "yaml",
			input:     "testdata/good.json",
			want:      true,
		},
		"stdin": {
			extension: "yaml",
			input:     "-",
			stdin:     `{"app": {"debug": false}}`,
			want:      false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdin = r
			w.WriteString(tc.stdin)
			w.Close()

			ctx := context.Background()

			c := config{}
			c.App.Debug = true

			assert.Equal(t, getRenderFiles(ctx, &c, tc.extension, tc.input) != nil, tc.err)
			assert.Equal(t, c.App.Debug, tc.want)
		})
	}

	os.Stdin = stdin
}
