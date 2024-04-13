package proto

import (
	"citizen/lib"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Load(srcAssets, dstAssets string) error {
	file, err := os.Create(filepath.Join(srcAssets, "proto.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	encode := json.NewEncoder(file)
	encode.SetIndent("", "\t")
	if err := encode.Encode(newProto()); err != nil {
		return err
	}

	if err := lib.CopyFile(filepath.Join(srcAssets, "proto.json"), filepath.Join(dstAssets, "proto.json")); err != nil {
		return err
	}

	// 对物品图集进行切分
	if err := cutPNG(srcAssets, dstAssets, "items"); err != nil {
		return err
	}

	return nil
}

func cutPNG(srcAssets, dstAssets string, dir string) error {
	if err := lib.Clean(filepath.Join(dstAssets, "images", dir)); err != nil {
		return err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(dstAssets, "images", dir))

	if err := filepath.Walk(filepath.Join(srcAssets, dir), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		assetName := strings.TrimRight(strings.ToLower(filepath.Base(path)), ".png")
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("os.Open(%q) failure :: %s", path, err.Error())
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return fmt.Errorf("image.Decode(%q) failure :: %s", path, err.Error())
		}

		bounds := img.Bounds()
		for x := 0; x < bounds.Dx(); x = x + 32 {
			for y := 0; y < bounds.Dy(); y = y + 32 {
				// 空文件模版
				size32File, err := os.Open(filepath.Join(srcAssets, "size32.png"))
				if err != nil {
					return fmt.Errorf("os.Open(%q) failure :: %s", filepath.Join(srcAssets, "size32.png"), err.Error())
				}

				// 读取Image对象
				size32Image, _, err := image.Decode(size32File)
				if err != nil {
					return fmt.Errorf("png.Decode(%q) failure :: %s", filepath.Join(srcAssets, "size32.png"), err.Error())
				}

				for x0 := 0; x0 < 32; x0++ {
					for y0 := 0; y0 < 32; y0++ {
						r, g, b, a := img.(*image.NRGBA).At(x+x0, y+y0).RGBA()
						newRGBA := color.RGBA{
							R: uint8(r >> 8),
							G: uint8(g >> 8),
							B: uint8(b >> 8),
							A: uint8(a >> 8),
						}

						size32Image.(*image.NRGBA).Set(x0, y0, newRGBA)
					}
				}

				// 写入文件
				newFileName := filepath.Join(dstAssets, "images", dir, fmt.Sprintf("%s_%d_%d.png", assetName, 1+y/32, 1+x/32))
				newFile, err := os.Create(newFileName)
				if err != nil {
					return fmt.Errorf("os.Create(%q) failure :: %s", newFileName, err.Error())
				}

				if err := png.Encode(newFile, size32Image); err != nil {
					return fmt.Errorf("png.Encode(%q) failure :: %s", newFileName, err.Error())
				}

				if err := newFile.Close(); err != nil {
					return fmt.Errorf("newFile.Close failure :: %s", err.Error())
				}

				if err := size32File.Close(); err != nil {
					return fmt.Errorf("newFile.Close failure :: %s", err.Error())
				}
			}
		}

		log.Printf("按照32*32尺寸将图集 %s/%s 切分成PNG图片 ... \n", dir, assetName)
		return nil
	}); err != nil {
		return err
	}

	return nil
}
