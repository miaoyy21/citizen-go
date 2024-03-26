package tools

import (
	"github.com/toy80/audio/vorbis"
	"github.com/toy80/audio/wav"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func OggToWav(srcAssets string) error {

	// 将ogg转为wav文件
	ogs := make([]string, 0)
	if err := filepath.Walk(filepath.Join(srcAssets, "audio"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(info.Name(), ".ogg") {
			ogs = append(ogs, path)
		}

		return nil
	}); err != nil {
		return err
	}

	for _, og := range ogs {
		ogFile, err := vorbis.Open(og)
		if err != nil {
			return err
		}

		ext := filepath.Ext(og)
		newFileName := strings.ReplaceAll(og, ext, ".wav")

		if err := wav.WriteFile(newFileName, ogFile); err != nil {
			return err
		}

		ogFile.Close()

		if err := os.Remove(og); err != nil {
			return err
		}

		log.Printf("完成音频文件%q转为%q ...\n", filepath.Base(og), filepath.Base(newFileName))
	}

	return nil
}
