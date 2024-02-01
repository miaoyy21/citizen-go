package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Mp4ToGif(root string) error {
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
	//if err := Mp4ToGif("assets"); err != nil {
	//	log.Fatalf("Fatal ERROR :: %s \n", err.Error())
	//}

	in, err := os.Open("assets/001.gif")
	if err != nil {
		log.Fatalf("os.Open ERROR :: %s \n", err.Error())
	}
	defer in.Close()

	g, err := gif.DecodeAll(in)
	if err != nil {
		log.Fatalf("gif.DecodeAll ERROR :: %s \n", err.Error())
	}

	log.Printf("Image Size => %d, Delay Size => %d, Loop Count => %d, Disposal Size => %d, Background Index => %#v \n", len(g.Image), len(g.Delay), g.LoopCount, len(g.Disposal), g.BackgroundIndex)
	log.Printf("Config Color Model => %#v \n", g.Config.ColorModel)
	log.Printf("Config Width => %d, Config Height => %d\n", g.Config.Width, g.Config.Height)

	img := g.Image[0]
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	log.Printf("Frame Width => %d, Frame Height => %d \n", w, h)
	log.Printf("Pix Size => %d, Palette Size => %d \n", len(img.Pix), len(img.Palette))

	skipped := NewSkippedColor(0xb4b4aaff, 0xd8d8aaff, 0xb4b4ffff)

	wImg := image.NewRGBA(image.Rect(0, 0, w, h))
	//wImg := image.NewRGBA64(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := img.At(x, y)

			//v := wImg.ColorModel().Convert(color.NRGBA64{R: uint16(r0), G: uint16(g0), B: uint16(b0), A: uint16(0)})
			//rr, gg, bb, aa := v.RGBA()
			//wImg.Set(x, y, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})

			if skipped.IsIn(c) {
				wImg.Set(x, y, color.RGBA{A: uint8(0)})

				continue
			}

			wImg.Set(x, y, c)
		}
	}

	out, err := os.Create("out.png")
	if err != nil {
		log.Fatalf("os.Create ERROR :: %s \n", err.Error())
	}
	defer out.Close()

	if err := png.Encode(out, wImg); err != nil {
		log.Fatalf("png.Encode ERROR :: %s \n", err.Error())
	}

	log.Println("Process Finished ...")
}

type SkippedColor map[uint32]struct{}

func NewSkippedColor(rgba ...uint32) SkippedColor {
	skipped := make(map[uint32]struct{}, len(rgba))

	//	R: V >> 24
	//	G: V << 8 >> 24
	//	B: V << 16 >> 24
	//	A: V << 24 >> 24
	for _, c32 := range rgba {
		skipped[c32] = struct{}{}
	}

	return skipped
}

func (skipped SkippedColor) IsIn(c64 color.Color) bool {
	r1, g1, b1, a1 := c64.RGBA()

	rgba := r1>>8<<24 | g1>>8<<16 | b1>>8<<8 | a1>>8
	if _, ok := skipped[rgba]; ok {
		return true
	}

	return false
}
