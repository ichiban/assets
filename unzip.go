package assets

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// Unzip locates resources by unzipping the executable.
// If it detects an appended zip file in the executable binary, it extracts resources in a temporary directory and returns the path.
// Otherwise, it fails.
func Unzip() Strategy {
	return &unzipStrategy{}
}

type unzipStrategy struct {
	tempDir string
}

func (s *unzipStrategy) Path() (string, error) {
	exec, err := os.Executable()
	if err != nil {
		return "", errors.Wrap(err, "failed to get executable")
	}

	r, err := zip.OpenReader(exec)
	if err != nil {
		return "", errors.Wrap(err, "failed to get zip reader")
	}

	tempDir, err := ioutil.TempDir("", filepath.Base(exec))
	if err != nil {
		return "", errors.Wrapf(err, "failed to open template directory: %s", filepath.Base(exec))
	}
	s.tempDir = tempDir

	for _, f := range r.File {
		if err := extract(f, tempDir); err != nil {
			return "", errors.Wrapf(err, "failed to extract file: %s", f.Name)
		}
	}

	return tempDir, nil
}

func (s *unzipStrategy) Close() error {
	if s.tempDir == "" {
		return nil
	}
	return os.RemoveAll(s.tempDir)
}

func extract(f *zip.File, dir string) error {
	if strings.Contains(f.Name, "..") {
		// Zip Slip!
		return fmt.Errorf("file path '%s' contains '..'", f.Name)
	}

	path := filepath.Join(dir, filepath.Clean(f.Name))

	if f.Mode().IsDir() {
		if err := os.MkdirAll(path, f.Mode()); err != nil {
			return errors.Wrapf(err, "failed to make a directory '%s'", path)
		}
		return nil
	}

	r, err := f.Open()
	if err != nil {
		return errors.Wrapf(err, "failed to open an archived file '%s'", f.Name)
	}

	tf, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return errors.Wrapf(err, "failed to open the target file '%s'", path)
	}
	defer func() {
		_ = tf.Close()
	}()

	if _, err := io.Copy(tf, r); err != nil {
		return errors.Wrap(err, "failed to copy")
	}

	return nil
}
