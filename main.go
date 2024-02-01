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

	srcImg := g.Image[0]
	w, h := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()
	log.Printf("Frame Width => %d, Frame Height => %d \n", w, h)
	log.Printf("Pix Size => %d, Palette Size => %d \n", len(srcImg.Pix), len(srcImg.Palette))

	replaced := NewReplacedColor(map[color.Color][]uint32{
		color.RGBA{A: 0}: {0xb4b4aaff, 0xd8d8aaff, 0xb4b4ffff},
	})

	dstImg := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			src := srcImg.At(x, y)
			dstImg.Set(x, y, replaced.Replace(src))
		}
	}

	out, err := os.Create("out.png")
	if err != nil {
		log.Fatalf("os.Create ERROR :: %s \n", err.Error())
	}
	defer out.Close()

	if err := png.Encode(out, dstImg); err != nil {
		log.Fatalf("png.Encode ERROR :: %s \n", err.Error())
	}

	log.Println("Process Finished ...")
}

type ReplacedColor map[uint32]color.Color

func NewReplacedColor(rgba map[color.Color][]uint32) ReplacedColor {
	skipped := make(map[uint32]color.Color)

	//	R: V >> 24
	//	G: V << 8 >> 24
	//	B: V << 16 >> 24
	//	A: V << 24 >> 24
	for c, c32s := range rgba {
		for _, c32 := range c32s {
			skipped[c32] = c
		}
	}

	return skipped
}

func (replaced ReplacedColor) Replace(src color.Color) color.Color {
	r, g, b, a := src.RGBA()

	rgba := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
	if dst, ok := replaced[rgba]; ok {
		return dst
	}

	return src
}
