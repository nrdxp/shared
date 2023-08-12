package cli

import (
	"context"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/config"
	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
)

type C struct {
	CLI   Config
	Hide  msg
	Other msg
	Show  msg
}

func (c *C) CLIConfig() *Config {
	return &c.CLI
}

func (c *C) Parse(ctx context.Context, configArgs, paths string) errs.Err {
	return config.Parse(ctx, c, "cli", "", configArgs, paths)
}

type msg struct {
	Message string `json:"Message,omitempty"` //nolint:tagliatelle
}

func TestPrint(t *testing.T) {
	ctx := context.Background()

	c := C{
		Hide: msg{
			Message: "Hello",
		},
		Show: msg{
			Message: "World",
		},
	}

	tests := map[string]struct {
		input  bool
		output string
	}{
		"json": {
			input: true,
			output: `{
  "CLI": {
    "debug": false,
    "noColor": false,
    "outputJSON": true
  },
  "Hide": {
    "Message": "Hello"
  },
  "Show": {
    "Message": "World"
  }
}
`,
		},
		"yaml": {
			output: `CLI:
  debug: false
  noColor: false
  outputJSON: false
Hide:
  Message: Hello
Show:
  Message: World
`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			logger.SetStd()

			c.CLI.OutputJSON = tc.input

			assert.Equal(t, printConfig(ctx, App[*C]{
				Config: &c,
				HideConfigFields: []string{
					"Hide",
					"Other",
				},
			}), nil)

			assert.Equal(t, logger.ReadStd(), tc.output)
		})
	}
}
