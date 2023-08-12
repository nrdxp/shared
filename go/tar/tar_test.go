package tar

import (
	"bytes"
	"compress/gzip"
	"context"
	"os"
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestTar(t *testing.T) {
	logger.UseTestLogger(t)

	ctx := context.Background()

	b := bytes.Buffer{}
	w := gzip.NewWriter(&b)

	assert.Equal(t, Create(ctx, "testdata", w), nil)
	assert.Equal(t, len(b.Bytes()), 10)

	w.Close()

	r, _ := gzip.NewReader(bytes.NewBuffer(b.Bytes()))

	f, err := List(ctx, r)
	assert.Equal(t, err, nil)
	assert.Equal(t, f, []string{
		"/dir1",
		"/dir1/file1.txt",
		"/dir2",
		"/dir2/file2.txt",
	})

	r, _ = gzip.NewReader(bytes.NewBuffer(b.Bytes()))

	os.MkdirAll("testdata/restore", 0755)
	assert.Equal(t, Extract(ctx, r, "testdata/restore"), nil)

	i, err := os.Stat("testdata/restore/dir2")
	assert.Equal(t, err, nil)
	assert.Equal(t, i.Mode().Perm(), 0755)

	out, err := os.ReadFile("testdata/restore/dir2/file2.txt")
	assert.Equal(t, err, nil)
	assert.Equal(t, string(out), "Hello world\n")

	os.RemoveAll("testdata/restore")
}
