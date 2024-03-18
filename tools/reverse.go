package tools

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

// 根据原始文件，产生向左和向右的文件
func reverse(originalFileName string, symbol Symbol, dstRoot string) error {
	fileName := filepath.Base(originalFileName)

	leftFileName := filepath.Join(dstRoot, fmt.Sprintf("%s_%s_%s", string(symbol), DirectionLeft, fileName))
	rightFileName := filepath.Join(dstRoot, fmt.Sprintf("%s_%s_%s", string(symbol), DirectionRight, fileName))

	// 复制1份向右的文件
	if err := CopyFile(originalFileName, rightFileName); err != nil {
		return fmt.Errorf("reverse CopyFile [%s] : %s", originalFileName, err.Error())
	}

	rightFile, err := os.Open(rightFileName)
	if err != nil {
		return fmt.Errorf("reverse os.Open [%s] : %s", originalFileName, err.Error())
	}
	defer rightFile.Close()

	leftImage, _, err := image.Decode(rightFile)
	if err != nil {
		return fmt.Errorf("reverse image.Decode [%s] : %s", originalFileName, err.Error())
	}

	colors := make(map[int]map[int]color.Color)
	bounds := leftImage.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			xColors, ok := colors[x]
			if !ok {
				xColors = make(map[int]color.Color)
			}

			r, g, b, a := leftImage.(*image.NRGBA).At(x, y).RGBA()
			r0, g0, b0, a0 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
			xColors[y] = color.RGBA{R: r0, G: g0, B: b0, A: a0}
			colors[x] = xColors
		}
	}

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			leftImage.(*image.NRGBA).Set(bounds.Dx()-1-x, y, colors[x][y])
		}
	}

	// 向左
	leftFile, err := os.Create(leftFileName)
	if err != nil {
		return fmt.Errorf("reverse os.Create [%s] : %s", originalFileName, err.Error())
	}
	defer leftFile.Close()

	if err := png.Encode(leftFile, leftImage); err != nil {
		return fmt.Errorf("reverse png.Encode [%s] : %s", originalFileName, err.Error())
	}

	return nil
}

// 根据原始文件，产生向左和向右的文件
func onlyReverse(originalFileName string, dstRoot string) error {
	fileName := filepath.Base(originalFileName)
	leftFileName := filepath.Join(dstRoot, fileName)

	rightFile, err := os.Open(originalFileName)
	if err != nil {
		return fmt.Errorf("reverse os.Create [%s] : %s", originalFileName, err.Error())
	}
	defer rightFile.Close()

	leftImage, _, err := image.Decode(rightFile)
	if err != nil {
		return fmt.Errorf("reverse image.Decode [%s] : %s", originalFileName, err.Error())
	}

	colors := make(map[int]map[int]color.Color)
	bounds := leftImage.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			xColors, ok := colors[x]
			if !ok {
				xColors = make(map[int]color.Color)
			}

			r, g, b, a := leftImage.(*image.NRGBA).At(x, y).RGBA()
			r0, g0, b0, a0 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
			xColors[y] = color.RGBA{R: r0, G: g0, B: b0, A: a0}
			colors[x] = xColors
		}
	}

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			leftImage.(*image.NRGBA).Set(bounds.Dx()-1-x, y, colors[x][y])
		}
	}

	// 向左
	leftFile, err := os.Create(leftFileName)
	if err != nil {
		return fmt.Errorf("reverse os.Create [%s] : %s", originalFileName, err.Error())
	}
	defer leftFile.Close()

	if err := png.Encode(leftFile, leftImage); err != nil {
		return fmt.Errorf("reverse png.Encode [%s] : %s", originalFileName, err.Error())
	}

	return nil
}
