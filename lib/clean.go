package lib

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Clean(root string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		return nil
	})
}
