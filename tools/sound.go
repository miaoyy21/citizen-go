package tools

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Sound struct {
	Name     string        // 技能编号
	Category SoundCategory // 类型
	Sequence int           // 帧号
}

type SoundOut struct {
	Name     string
	Sequence int
	Category SoundCategoryOut
	Size     int
	Audios   []string
}

func RunSounds(srcAssets string, dstAssets string) error {

	sounds := make(map[string][]*Sound)
	if err := filepath.Walk(filepath.Join(srcAssets, "sound"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		sound, err := parseSound(path)
		if err != nil {
			return err
		}

		if sound != nil {
			ss, ok := sounds[sound.Name]
			if !ok {
				ss = make([]*Sound, 0)
			}

			ss = append(ss, sound)
			sounds[sound.Name] = ss
		}

		return nil
	}); err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(dstAssets, "sounds.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	for _, ss := range sounds {
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Sequence < ss[j].Sequence
		})
	}

	audios, err := parseAudio(srcAssets, dstAssets)
	if err != nil {
		return err
	}

	out := make(map[string][]*SoundOut)
	for name, ss := range sounds {
		var start int
		var category SoundCategoryOut

		for _, s := range ss {
			var size int

			switch s.Category {
			case SoundCategorySwingStart:
				category = SoundCategoryOutSwing
				start = s.Sequence
			case SoundCategorySwingEnd:
				size = s.Sequence - start + 1
			case SoundCategoryHandStart:
				category = SoundCategoryOutHand
				start = s.Sequence
			case SoundCategoryHandEnd:
				size = s.Sequence - start + 1
			case SoundCategoryFootStart:
				category = SoundCategoryOutFoot
				start = s.Sequence
			case SoundCategoryFootEnd:
				size = s.Sequence - start + 1
			}

			if size > 0 {
				sos, ok := out[name]
				if !ok {
					sos = make([]*SoundOut, 0)
				}

				as := make([]string, 0)
				for i := size; i > 0; i-- {
					k := fmt.Sprintf("%s_%d", category, i)

					if vs, ok := audios[k]; ok {
						as = append(as, vs...)
						break
					}
				}

				so := &SoundOut{
					Name:     name,
					Sequence: start,
					Category: category,
					Size:     size,
					Audios:   as,
				}

				sos = append(sos, so)
				if len(sos) == 0 {
					log.Panicf("Sound List is EMPTY")
				}
				out[name] = sos
			}
		}
	}

	if err := json.NewEncoder(file).Encode(out); err != nil {
		return err
	}
	log.Printf("发布动作音频文件至%q ... \n", filepath.Join(dstAssets, "sounds.json"))

	return nil
}

func parseSound(path string) (*Sound, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[%s] %s", path, err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("[%s] %s", path, err.Error())
	}

	bounds := img.Bounds()
	n0 := strings.Split(filepath.Base(path), ".")[0]
	n0s := strings.Split(n0, "_")
	if len(n0s) != 2 {
		return nil, fmt.Errorf("[%s] %s", path, "invalid format file name")
	}

	seq, err := strconv.Atoi(n0s[1])
	if err != nil {
		return nil, fmt.Errorf("[%s] %s", path, err.Error())
	}

	//识别矩形框，锚点为左下角
	sound := &Sound{
		Name:     n0s[0],
		Sequence: seq,
		Category: SoundCategorySwingStart,
	}

	for x := 0; x < bounds.Dx(); x++ {
		for y := bounds.Dy(); y >= 0; y-- {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}

			rgba := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
			if rgba == 0x000000ff {
				continue
			}

			switch rgba {
			case 0x800000ff:
				sound.Category = SoundCategoryHandStart
				return sound, nil
			case 0xff0000ff:
				sound.Category = SoundCategoryHandEnd
				return sound, nil
			case 0x800080ff:
				sound.Category = SoundCategoryFootStart
				return sound, nil
			case 0xff00ffff:
				sound.Category = SoundCategoryFootEnd
				return sound, nil
			case 0x008000ff:
				sound.Category = SoundCategorySwingStart
				return sound, nil
			case 0x00ff00ff:
				sound.Category = SoundCategorySwingEnd
				return sound, nil
			default:
				return nil, fmt.Errorf("[%s] unrecognize color %x at local point (%d,%d)", path, rgba, x, y)
			}
		}
	}

	return nil, nil
}
