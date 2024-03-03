package tools

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Clean(root string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		ns := strings.Split(strings.TrimRight(info.Name(), ".png"), "_")
		if len(ns) < 2 {
			return nil
		}

		if _, err := strconv.Atoi(ns[0]); err != nil {
			log.Printf("%s not remove \n", info.Name())
			return nil
		}

		if _, err := strconv.Atoi(ns[len(ns)-1]); err != nil {
			log.Printf("%s not remove \n", info.Name())
			return nil
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		log.Printf("%s Removed \n", info.Name())

		return nil
	})
}
