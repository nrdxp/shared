package logger

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/errs"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	ctx = SetAttribute(ctx, "key1", "value1")
	ctx = SetAttribute(ctx, "key2", "value2")

	tests := map[string]struct {
		debug      bool
		err        errs.Err
		noColor    bool
		outputWant string
		outputNot  string
	}{
		"error": {
			debug:      false,
			err:        errs.NewServerErr(errors.New("test")),
			outputWant: ColorRed + "[ERROR]",
		},
		"error no color": {
			debug:     false,
			err:       errs.NewServerErr(errors.New("test")),
			noColor:   true,
			outputNot: ColorRed,
		},
		"debug disabled": {
			debug:     false,
			err:       errs.NewClientBadRequestErr("test", errors.New("test")),
			outputNot: "DEBUG",
		},
		"debug": {
			debug:      true,
			err:        errs.NewClientBadRequestErr("test", errors.New("test")),
			outputWant: "DEBUG",
		},
		"none": {
			debug:     true,
			err:       errs.NewServerErr(),
			outputNot: "ERROR",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			SetStd()

			ctx = SetDebug(ctx, tc.debug)
			noColor = tc.noColor

			Log(ctx, tc.err)

			out := ReadStd()

			if tc.outputWant == "" {
				assert.Equal(t, strings.Contains(out, tc.outputNot), false)
			} else {
				assert.Contains(t, out, "logger.TestLog.func1")
				assert.Contains(t, out, "key2='value2'")
				assert.Contains(t, out, tc.outputWant)
			}
		})
	}
}
