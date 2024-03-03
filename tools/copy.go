package tools

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Copy(oRoot, tRoot string) error {
	return filepath.Walk(oRoot, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == "animations.json" {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		oFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer oFile.Close()

		tFile, err := os.Create(filepath.Join(tRoot, info.Name()))
		if err != nil {
			return err
		}
		defer tFile.Close()

		if _, err := io.Copy(tFile, oFile); err != nil {
			return err
		}

		log.Printf("%s Copyed \n", info.Name())
		return nil
	})
}
