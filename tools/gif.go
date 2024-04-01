package tools

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/fs"
	"path/filepath"
	"strings"
)

func Mp4SwfToGif(root string) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.EqualFold(filepath.Ext(path), ".mp4") {
			// 1280*720 = 16:9 = 320*180
			return ffmpeg.Input(path).
				Output(strings.ReplaceAll(path, ".mp4", ".gif"), ffmpeg.KwArgs{"s": "1920x1080", "r": "12"}).
				OverWriteOutput().ErrorToStdOut().Run()
		} else if strings.EqualFold(filepath.Ext(path), ".swf") {
			//return ffmpeg.Input(path).
			//	Output(strings.ReplaceAll(path, ".swf", ".gif"), ffmpeg.KwArgs{"s": "1920x1080", "r": "12"}).
			//	OverWriteOutput().ErrorToStdOut().Run()
		}

		return nil
	})
}
