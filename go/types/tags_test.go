package types

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestTagsParseSlice(t *testing.T) {
	tests := map[string]struct {
		input SliceString
		want  Tags
	}{
		"empty": {
			input: SliceString{},
			want:  Tags{},
		},
		"valid": {
			input: SliceString{
				"a tag With spaces",
				"Tag",
				"a tag With spaces",
				"tag",
			},
			want: Tags{
				"atagwithspaces",
				"tag",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var tags Tags

			tags.ParseSlice(tc.input)

			assert.Equal(t, tags, tc.want)
		})
	}
}
