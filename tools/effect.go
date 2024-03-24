package tools

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
)

func changeEffect(fileName string, effect Effect) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("changeEffect os.Open [%s] : %s", fileName, err.Error())
	}

	newImage, _, err := image.Decode(file)
	if err != nil {
		file.Close()
		return fmt.Errorf("changeEffect image.Decode [%s] : %s", fileName, err.Error())
	}
	file.Close()

	bounds := newImage.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r, g, b, a := newImage.(*image.NRGBA).At(x, y).RGBA()
			r0, g0, b0, a0 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)

			rgb := make([]uint8, 0, 3)
			rgb = append(rgb, r0, g0, b0)
			sort.Slice(rgb, func(i, j int) bool {
				return rgb[i] > rgb[j]
			})

			v1, v2, v3 := rgb[0], rgb[1], rgb[2]
			if effect == EffectRed {
				newImage.(*image.NRGBA).Set(x, y, color.RGBA{R: v1, G: v3, B: v2, A: a0})
			} else if effect == EffectGreen {
				newImage.(*image.NRGBA).Set(x, y, color.RGBA{R: v2, G: v1, B: v3, A: a0})
			} else if effect == EffectBlue {
				newImage.(*image.NRGBA).Set(x, y, color.RGBA{R: v3, G: v2, B: v1, A: a0})
			}
		}
	}

	// 文件重写
	newFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("changeEffect os.Create [%s] : %s", fileName, err.Error())
	}
	defer newFile.Close()

	if err := png.Encode(newFile, newImage); err != nil {
		return fmt.Errorf("changeEffect png.Encode [%s] : %s", fileName, err.Error())
	}

	return nil
}
