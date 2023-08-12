package template

import (
	"context"
	"net/http"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestLint(t *testing.T) {
	logger.UseTestLogger(t)

	ctx := context.Background()

	tmpl := `
hello {{ omfg }}
{{ len() }}
{{ len 1 }}
{{ if eq 1 0 }}
{{ .data 1 1 23j4k4j4kjasdf cxzbjsjdkafjds }}
{{ $hello := wtf }}
{{ $hello := "wtf" }}
{{ else }}
{{ $hello }}
`

	out, err := Lint(ctx, tmpl, nil)

	assert.HasErr(t, err, nil)
	assert.Equal(t, out, LintResults{
		2: []string{`function "omfg" not defined`},
		3: []string{`unexpected "(" in operand`},
		4: []string{`executing "render" at <len 1>: error calling len: len of type int`},
		5: []string{"missing {{ end }}"},
		6: []string{`bad number syntax: "23j"`},
		7: []string{`function "wtf" not defined`},
		9: []string{"unexpected {{else}}"},
	})

	out[0] = []string{"hello"}
	assert.Equal(t, out.Show("my/file"), []string{`my/file: hello`, `my/file:2: function "omfg" not defined`, `my/file:3: unexpected "(" in operand`, `my/file:4: executing "render" at <len 1>: error calling len: len of type int`, `my/file:5: missing {{ end }}`, `my/file:6: bad number syntax: "23j"`, `my/file:7: function "wtf" not defined`, `my/file:9: unexpected {{else}}`})

	out, _ = Lint(ctx, `{{ $contents := try .contents "" }}
{{ $mode := try .mode "" }}
{{ $path := try .path "" }}
{{ $remove := try .remove false }}

{{ if (try .test.mock false) }}
{{ $contents = "hello" }}
{{ $mode = "0644" }}
{{ $path = "/path" }}
{{ $remove = true }}
{{ end }}

- id: file {{ $path }}
  check: |
    [[ $(crc32 <(cat <<EOF
{{ $contents | indent 4 }}
    EOF
    )) == $(crc32 {{ $path }}) ]] && [[ $(stat -c "%a" {{ $path }}) == {{ if eq (len $mode) 4 }}{{ $mode }}{{ else }}0{{ $mode }}{{ end }} ]]
  change: |
    cat > {{ $path }} <<EOF
{{ $contents | indent 4 }}
    EOF
    {{ if $mode }}
    chmod {{ $mode }} {{ $path }}
    {{ end }}
{{ if $remove }}
  remove: rm {{ $path }}
{{ end }}
`, http.Client{})
	assert.Equal(t, len(out), 21)
}
