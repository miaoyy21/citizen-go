package tools

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Audio struct {
	name     string
	duration float64
	frames   float64
}

func RunAudio(srcAssets string) error {
	audios := make(map[string][]Audio, 0)

	if err := filepath.Walk(filepath.Join(srcAssets, "audio"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !strings.EqualFold(ext, ".wav") && !strings.EqualFold(ext, ".ogg") {
			return nil
		}

		fileName := strings.TrimPrefix(path, filepath.Join(srcAssets, "audio"))

		cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", path)
		if cmd.Err != nil {
			return cmd.Err
		}

		bs, err := cmd.Output()
		if err != nil {
			return err
		}

		duration, err := strconv.ParseFloat(string(bytes.TrimSpace(bs)), 10)
		if err != nil {
			return err
		}

		dir := strings.Split(fileName, "_")[0]
		audio, ok := audios[dir]
		if !ok {
			audio = make([]Audio, 0)
		}

		audio = append(audio, Audio{
			name:     filepath.Base(fileName),
			duration: duration,
			frames:   duration * 12.0,
		})
		audios[dir] = audio

		return nil
	}); err != nil {
		return err
	}

	for _, audio := range audios {
		sort.Slice(audio, func(i, j int) bool {
			return audio[i].frames < audio[j].frames
		})
	}

	for dir, audio := range audios {
		for _, sub := range audio {
			log.Printf("%-6s %-20s %9.4f %6.3f", dir, sub.name, sub.duration, sub.frames)
		}
		fmt.Println()
	}

	return nil
}

func parseAudio(srcAssets string, dstAssets string) (map[string][]string, error) {
	sounds := make(map[string][]string)

	// 2. 清空目标文件夹
	if err := clean(filepath.Join(dstAssets, "audio")); err != nil {
		return nil, err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(dstAssets, "audio"))

	// 1. 检索资源文件
	if err := filepath.Walk(filepath.Join(srcAssets, "audio"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !strings.EqualFold(ext, ".wav") && !strings.EqualFold(ext, ".ogg") {
			return nil
		}

		fileName := filepath.Base(path)

		n3s := strings.Split(strings.Split(fileName, ".")[0], "_")
		if strings.HasPrefix(fileName, "Blow") || strings.HasPrefix(fileName, "Swing") {
			n12 := fmt.Sprintf("%s_%s", n3s[0], n3s[1])
			ss, ok := sounds[n12]
			if !ok {
				ss = make([]string, 0)
			}

			ss = append(ss, fileName)
			sounds[n12] = ss
		}

		if err := CopyFile(path, filepath.Join(dstAssets, "audio", info.Name())); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	log.Printf("发布音效%q至目标目录  ... \n", filepath.Join(dstAssets, "audio"))

	return sounds, nil
}
