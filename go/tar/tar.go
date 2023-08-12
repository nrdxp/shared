// Package tar is used to create and extract tar files.
package tar

import (
	"archive/tar"
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/candiddev/shared/go/errs"
	"github.com/candiddev/shared/go/logger"
)

var ErrCreate = errors.New("error creating archive")
var ErrExtract = errors.New("error extracting archive")
var ErrList = errors.New("error listing archive")

// Create performs a tar of the source path to the destination writer.
func Create(ctx context.Context, source string, destination io.Writer) error {
	ctx = logger.Trace(ctx)

	t := tar.NewWriter(destination)

	defer t.Close()

	if err := filepath.Walk(source, func(path string, info fs.FileInfo, _err error) error {
		if path == source {
			return nil
		}

		header, err := tar.FileInfoHeader(info, path)
		if err != nil {
			return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
		}

		header.Name = strings.TrimPrefix(path, source)

		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			header.Linkname, err = os.Readlink(path)
			if err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
			}
		}

		if err := t.WriteHeader(header); err != nil {
			return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
		}

		if info.Mode().IsRegular() {
			f, err := os.Open(path)
			if err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
			}

			if _, err := io.Copy(t, f); err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
			}
		}

		return nil
	}); err != nil {
		return logger.Log(ctx, errs.NewServerErr(ErrCreate, err))
	}

	return logger.Log(ctx, nil)
}

// Extract uses an io.Reader to extract a tar archive to destination.
func Extract(ctx context.Context, source io.Reader, destination string) error {
	ctx = logger.Trace(ctx)

	t := tar.NewReader(source)

	for {
		header, err := t.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return logger.Log(ctx, errs.NewServerErr(ErrExtract, err))
		}

		target := filepath.Join(destination, header.Name) //nolint:gosec

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(target, fs.FileMode(header.Mode)); err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrExtract, err))
			}

		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrExtract, err))
			}

			if _, err := io.Copy(f, t); err != nil { //nolint:gosec
				return logger.Log(ctx, errs.NewServerErr(ErrExtract, err))
			}

			f.Close()
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return logger.Log(ctx, errs.NewServerErr(ErrExtract, err))
			}
		}
	}

	return nil
}

// List uses a source io.Reader to list the contents of the archive.
func List(ctx context.Context, source io.Reader) ([]string, error) {
	ctx = logger.Trace(ctx)

	t := tar.NewReader(source)

	files := []string{}

	for {
		header, err := t.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, logger.Log(ctx, errs.NewServerErr(ErrList, err))
		}

		files = append(files, header.Name)
	}

	return files, nil
}
