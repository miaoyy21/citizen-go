package tools

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CopyDirectory(srcRoot, dstRoot string) error {
	return filepath.Walk(srcRoot, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == "animations.json" {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		if err := CopyFile(path, filepath.Join(dstRoot, info.Name())); err != nil {
			return err
		}

		return nil
	})
}

func CopyFile(srcFileName, dstFileName string) error {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}
