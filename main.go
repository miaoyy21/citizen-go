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

	c1 := NewUint32Color(0xb4b4aaff)
	c2 := NewUint32Color(0xd8d8aaff)
	c3 := NewUint32Color(0xb4b4ffff)

	wImg := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := img.At(x, y)

			r0, g0, b0, a0 := c.RGBA()
			if c1.Equal(c) || c2.Equal(c) || c3.Equal(c) {

				//rr, gg, bb, aa := wImg.ColorModel().Convert(color.NRGBA{A: uint8(0)}).RGBA()
				////rr, gg, bb, aa := color.NRGBA{A: uint8(0)}.RGBA()
				//log.Printf("(%d,%d) => (%x,%x,%x,%x) \n", x, y, rr, gg, bb, aa)
				//wImg.Set(x, y, color.RGBA{A: uint8(aa)})
				//
				v := wImg.ColorModel().Convert(color.NRGBA64{R: uint16(r0), G: uint16(g0), B: uint16(b0), A: uint16(0)})
				rr, gg, bb, aa := v.RGBA()
				wImg.Set(x, y, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})

				log.Printf("(%d,%d) => Continue ... \n", x, y)
				continue
			}

			log.Printf("(%d,%d) => (%x,%x,%x,%x) \n", x, y, r0, g0, b0, a0)
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

type Uint32Color struct {
	rgba color.RGBA
}

func NewUint32Color(rgba uint32) Uint32Color {
	return Uint32Color{
		rgba: color.RGBA{
			R: uint8(rgba >> 24),
			G: uint8((rgba << 8) >> 24),
			B: uint8((rgba << 16) >> 24),
			A: uint8(rgba << 24 >> 24),
		},
	}
}

func (c32 Uint32Color) Equal(c64 color.Color) bool {
	r1, g1, b1, a1 := c64.RGBA()
	r0, g0, b0, a0 := c32.rgba.RGBA()

	return r1 == r0 && g1 == g0 && b1 == b0 && a1 == a0
}
