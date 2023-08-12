package config

import (
	"context"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestLint(t *testing.T) {
	logger.UseTestLogger(t)

	ctx := context.Background()

	tests := map[string]struct {
		extension string
		wantLint  LintResults
		wantOut   LintOutputs[*config]
		wantShow  []string
	}{
		"testdata/good.json": {
			extension: "json",
			wantOut: LintOutputs[*config]{
				"testdata/good.json": {
					Config: &config{
						App: configApp{
							Lists: []string{
								"hello",
								"world",
							},
						},
					},
					Text: `{
  "app": {
    "debug": false,
    "lists": ["hello", "world"]
  }
}
`,
				},
			},
			wantLint: LintResults{},
			wantShow: []string{},
		},
		"testdata/templates/bad": {
			extension: "yaml",
			wantLint: LintResults{
				"testdata/templates/bad/bad1.yaml": {
					0: []string{`error rendering config: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type config.config`},
				},
				"testdata/templates/bad/bad2.yaml": {
					0: []string{`error rendering config: error parsing template: template: render:1: unexpected {{else}}`},
					1: []string{`unexpected {{else}}`},
					2: []string{`function "oops" not defined`},
					3: []string{"unexpected {{end}}"},
				},
			},
			wantShow: []string{
				`testdata/templates/bad/bad1.yaml: error rendering config: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type config.config`,
				`testdata/templates/bad/bad2.yaml: error rendering config: error parsing template: template: render:1: unexpected {{else}}`,
				`testdata/templates/bad/bad2.yaml:1: unexpected {{else}}`,
				`testdata/templates/bad/bad2.yaml:2: function "oops" not defined`,
				`testdata/templates/bad/bad2.yaml:3: unexpected {{end}}`,
			},
		},
		"testdata/invalid.json": {
			wantLint: LintResults{
				"testdata/invalid.json": {
					0: []string{`error rendering config: error unmarshaling JSON: while decoding JSON: json: unknown field "this is not valid json"`},
				},
			},
			wantShow: []string{
				`testdata/invalid.json: error rendering config: error unmarshaling JSON: while decoding JSON: json: unknown field "this is not valid json"`,
			},
		},
		"testdata/invalid.yaml": {
			wantLint: LintResults{
				"testdata/invalid.yaml": {
					0: []string{"error rendering config: error converting YAML to JSON: yaml: did not find expected key"},
				},
			},
			wantShow: []string{
				`testdata/invalid.yaml: error rendering config: error converting YAML to JSON: yaml: did not find expected key`,
			},
		},
		"testdata/good.yaml": {
			extension: "Json",
			wantLint:  LintResults{},
			wantShow:  []string{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			outRes, outOut, err := Lint(ctx, tc.extension, name, config{}, func() *config {
				return &config{}
			})
			assert.HasErr(t, err, nil)
			assert.Equal(t, outRes, tc.wantLint)
			assert.Equal(t, outRes.Show(), tc.wantShow)

			if tc.wantOut != nil {
				assert.Equal(t, outOut, tc.wantOut)
			}
		})
	}
}
