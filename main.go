package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func mp4ToGif(root string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.EqualFold(filepath.Ext(path), ".mp4") {
			return nil
		}

		return ffmpeg.Input(path).
			Output(strings.ReplaceAll(path, ".mp4", ".gif"), ffmpeg.KwArgs{"s": "1280x720", "r": "12"}).
			OverWriteOutput().ErrorToStdOut().Run()
	})
}

func main() {
	log.Println("Process Started ...")
	if err := mp4ToGif("assets"); err != nil {
		log.Fatalf("Fatal ERROR :: %s \n", err.Error())
	}

	log.Println("Process Finished ...")
}
