package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/color"
	"io/fs"
	"log"
	"math"
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

	if err := Mp4ToGif("assets"); err != nil {
		log.Fatalf("Fatal ERROR :: %s \n", err.Error())
	}

	//in, err := os.Open("assets/001.gif")
	//if err != nil {
	//	log.Fatalf("os.Open ERROR :: %s \n", err.Error())
	//}
	//defer in.Close()
	//
	//g, err := gif.DecodeAll(in)
	//if err != nil {
	//	log.Fatalf("gif.DecodeAll ERROR :: %s \n", err.Error())
	//}
	//
	//log.Printf("Image Size => %d, Delay Size => %d, Loop Count => %d, Disposal Size => %d, Background Index => %#v \n", len(g.Image), len(g.Delay), g.LoopCount, len(g.Disposal), g.BackgroundIndex)
	//log.Printf("Config Color Model => %#v \n", g.Config.ColorModel)
	//log.Printf("Config Width => %d, Config Height => %d\n", g.Config.Width, g.Config.Height)
	//
	//srcImg := g.Image[8]
	//w, h := srcImg.Bounds().Dx(), srcImg.Bounds().Dy()
	//log.Printf("Frame Width => %d, Frame Height => %d \n", w, h)
	//log.Printf("Pix Size => %d, Palette Size => %d \n", len(srcImg.Pix), len(srcImg.Palette))
	//
	//replaced := NewReplacedColor(map[color.RGBA][]uint32{
	//	color.RGBA{A: 0}: {
	//		0xb4b4aaff, 0xd8d8aaff, 0xb4b4ffff,
	//		0x9090aaff, 0xd8d8ffff, 0xd8d8ffff,
	//		0x909055ff, 0xb4b455ff, 0x6c6caaff,
	//	},
	//})
	//
	//dstImg := image.NewRGBA(image.Rect(0, 0, w, h))
	//for x := 0; x < w; x++ {
	//	for y := 0; y < h; y++ {
	//		src := srcImg.At(x, y)
	//
	//		dst, ok := replaced.Replace(src)
	//		if ok {
	//			dstImg.Set(x, y, dst)
	//			continue
	//		}
	//
	//		dstImg.Set(x, y, dst)
	//	}
	//}
	//
	//// 消除锯齿
	//clear1(dstImg)
	//
	//// 消除噪音
	//clear2(dstImg)
	//
	//// PNG格式输出到文件
	//out, err := os.Create("out.png")
	//if err != nil {
	//	log.Fatalf("os.Create ERROR :: %s \n", err.Error())
	//}
	//defer out.Close()
	//
	//if err := png.Encode(out, dstImg); err != nil {
	//	log.Fatalf("png.Encode ERROR :: %s \n", err.Error())
	//}

	log.Println("Process Finished ...")
}

type ReplacedColor map[uint32]color.RGBA

func NewReplacedColor(rgba map[color.RGBA][]uint32) ReplacedColor {
	skipped := make(map[uint32]color.RGBA)

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

func (replaced ReplacedColor) Replace(src color.Color) (color.RGBA, bool) {
	r, g, b, a := src.RGBA()

	rgba := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
	if dst, ok := replaced[rgba]; ok {
		return dst, true
	}

	return color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}, false
}

func clear1(dstImg *image.RGBA) {
	for x := 1; x < dstImg.Rect.Dx()-1; x++ {
		for y := 1; y < dstImg.Rect.Dy()-1; y++ {
			src := dstImg.At(x, y)

			_, _, _, a := src.RGBA()
			if a == 0xffff {
				dstImg.Set(x, y, color.Black)
				continue
			}

			var sum int
			for x0 := x - 1; x0 <= x+1; x0++ {
				for y0 := y - 1; y0 <= y+1; y0++ {
					if x0 == x && y0 == y {
						continue
					}

					_, _, _, a0 := dstImg.At(x0, y0).RGBA()
					if a0 == 0xffff {
						sum += 1
					}
				}
			}

			if sum >= 5 {
				dstImg.Set(x, y, color.Black)
			}
		}
	}
}

func clear2(dstImg *image.RGBA) {

	for x := 1; x < dstImg.Rect.Dx()-1; x++ {
		for y := 1; y < dstImg.Rect.Dy()-1; y++ {
			src := dstImg.At(x, y)

			_, _, _, a := src.RGBA()
			if a != 0xffff {
				continue
			}

			for i := 1; i <= 8; i++ {
				var sum int
				for x0 := x - i; x0 <= x+i; x0++ {
					for y0 := y - i; y0 <= y+i; y0++ {
						if x0 == x && y0 == y {
							continue
						}

						_, _, _, a0 := dstImg.At(x0, y0).RGBA()
						if a0 == 0xffff {
							sum += 1
						}
					}
				}

				if sum <= int(math.Floor(0.25*float64(i*i*4))) {
					dstImg.Set(x, y, color.RGBA{A: 0})
				}
			}
		}
	}
}
