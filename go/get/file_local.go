package get

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fileLocal(_ context.Context, src string, dst io.Writer) (time.Time, error) {
	src = strings.TrimPrefix(src, "file:/")
	if strings.HasPrefix(src, "~") {
		dir, err := os.UserHomeDir()
		if err != nil {
			return time.Time{}, fmt.Errorf("error getting homedir: %w", err)
		}

		src = filepath.Join(dir, src[2:])
	}

	f, err := os.Open(src)
	if err != nil {
		return time.Time{}, fmt.Errorf("error opening src: %w", err)
	}

	s, err := f.Stat()
	if err != nil {
		return time.Time{}, fmt.Errorf("error getting stats for src: %w", err)
	}

	if dst != nil {
		if _, err := io.Copy(dst, f); err != nil {
			return time.Time{}, fmt.Errorf("error reading src: %w", err)
		}
	}

	return s.ModTime(), nil
}
