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

func RunAudio(dstAssets string) error {
	audios := make(map[string][]Audio, 0)

	if err := filepath.Walk(dstAssets, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileName := strings.TrimLeft(path, dstAssets)

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

		dir := filepath.Dir(fileName)
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
