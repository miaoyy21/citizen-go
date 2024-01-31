package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image/color"
	"image/gif"
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

	file, err := os.Open("assets/001.gif")
	if err != nil {
		log.Fatalf("os.Open ERROR :: %s \n", err.Error())
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		log.Fatalf("gif.Decode ERROR :: %s \n", err.Error())
	}

	log.Printf("Image Size => %d, Delay Size => %d, Loop Count => %d, Disposal Size => %d, Background Index => %#v \n", len(g.Image), len(g.Delay), g.LoopCount, len(g.Disposal), g.BackgroundIndex)
	log.Printf("Config Color Model => %#v \n", g.Config.ColorModel)
	log.Printf("Config Width => %d, Config Height => %d\n", g.Config.Width, g.Config.Height)

	img := g.Image[0]
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	log.Printf("Frame Width => %d, Frame Height => %d \n", w, h)
	log.Printf("Pix Size => %d, Palette Size => %d \n", len(img.Pix), len(img.Palette))

	rgba1 := NewUintColor(0xb4b4aa, 0xff)
	rgba2 := NewUintColor(0xd8d8aa, 0xff)
	rgba3 := NewUintColor(0xb4b4ff, 0xff)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rgba64 := img.At(x, y)

			if rgba1.Equal(rgba64) || rgba2.Equal(rgba64) || rgba3.Equal(rgba64) {
				continue
			}

			r0, g0, b0, a0 := rgba64.RGBA()

			log.Printf("(%d,%d) => (%x,%x,%x,%x) \n", x, y, r0, g0, b0, a0)
		}
	}

	log.Println("Process Finished ...")
}

type Uint32Color struct {
	rgba color.RGBA
}

func NewUintColor(rgb uint32, a uint8) Uint32Color {
	if rgb >= 0xffffff {
		log.Panicf("unexpect RGB Value %d \n", rgb)
	}

	log.Printf("G %x => %x \n", rgb, rgb&0x00ff00)
	v := Uint32Color{
		rgba: color.RGBA{
			R: uint8(rgb / 0xffff),
			G: uint8((rgb & 0x00ff00) / 0xff),
			B: uint8(rgb % 0xffff00),
			A: a,
		},
	}
	log.Printf("%x %x {%x,%x,%x,%x} \n", rgb, a, v.rgba.R, v.rgba.G, v.rgba.B, v.rgba.A)

	return v
}

func (c32 Uint32Color) Equal(c64 color.Color) bool {
	r1, g1, b1, a1 := c64.RGBA()
	r0, g0, b0, a0 := c32.rgba.RGBA()

	return r1 == r0 && g1 == g0 && b1 == b0 && a1 == a0
}
