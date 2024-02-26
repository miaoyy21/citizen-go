package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Animation struct {
	Name   string
	Width  int
	Height int

	IsLand   bool        // 是否在地面
	Position image.Point // 角色所在的粗略中心位置

	ExposeHead []image.Rectangle // 头（可被他人攻击）
	ExposeBody []image.Rectangle // 体（可被他人攻击）
	ExposeHand []image.Rectangle // 手（可被他人攻击）
	ExposeFoot []image.Rectangle // 脚（可被他人攻击）

	AttackHead []image.Rectangle // 头（可攻击他人）
	AttackBody []image.Rectangle // 体（可攻击他人）
	AttackHand []image.Rectangle // 手（可攻击他人）
	AttackFoot []image.Rectangle // 脚（可攻击他人）
}

// 计算角色所在的粗略中心位置
func (animation *Animation) setPosition() {
	min, max := image.Point{X: 10000, Y: 10000}, image.Point{}

	rects := make([]image.Rectangle, 0, len(animation.ExposeBody)+len(animation.AttackBody))
	rects = append(rects, animation.ExposeBody...)
	rects = append(rects, animation.AttackBody...)

	for _, rect := range rects {
		if rect.Min.X < min.X {
			min.X = rect.Min.X
		}

		if rect.Min.Y < min.Y {
			min.Y = rect.Min.Y
		}

		if rect.Max.X > max.X {
			max.X = rect.Max.X
		}

		if rect.Max.Y > max.Y {
			max.Y = rect.Max.Y
		}
	}

	animation.Position = image.Point{
		X: int(math.Ceil(float64(min.X+max.X) / 2)),
		Y: int(math.Ceil(float64(min.Y+max.Y) / 2)),
	}
}

func Walk(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !strings.HasSuffix(info.Name(), ".png") {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	log.Printf("文件名 %12s  尺寸 (W)%3d - (H)%-5d \n", info.Name(), bounds.Dx(), bounds.Dy())

	//识别矩形框，锚点为左下角
	animation := &Animation{
		Name:   strings.Split(info.Name(), ".")[0],
		Width:  bounds.Dx(),
		Height: bounds.Dy(),

		ExposeHead: make([]image.Rectangle, 0),
		ExposeBody: make([]image.Rectangle, 0),
		ExposeHand: make([]image.Rectangle, 0),
		ExposeFoot: make([]image.Rectangle, 0),

		AttackHead: make([]image.Rectangle, 0),
		AttackBody: make([]image.Rectangle, 0),
		AttackHand: make([]image.Rectangle, 0),
		AttackFoot: make([]image.Rectangle, 0),
	}

	skipped, isLand := make(map[string]struct{}), false
	for x := 0; x < bounds.Dx(); x++ {
		for y := bounds.Dy(); y >= 0; y-- {
			dx, dy := x, bounds.Dy()-1-y // 转换后的坐标系
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}

			rgba := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
			if rgba == 0x000000ff {
				// 正常的地面距离底部5像素
				if dy <= 7 {
					isLand = true
				}

				continue
			}

			if _, ok := skipped[fmt.Sprintf("%d_%d", dx, dy)]; ok {
				continue
			}

			rectangle := findRectangle(img, x, y, dx, dy, rgba)
			switch rgba {
			case 0x008000ff:
				animation.ExposeHead = append(animation.ExposeHead, rectangle)
			case 0x00ff00ff:
				animation.ExposeBody = append(animation.ExposeBody, rectangle)
			case 0x000080ff:
				animation.ExposeHand = append(animation.ExposeHand, rectangle)
			case 0x0000ffff:
				animation.ExposeFoot = append(animation.ExposeFoot, rectangle)
			case 0xff8000ff:
				animation.AttackHead = append(animation.AttackHead, rectangle)
			case 0xffff00ff:
				animation.AttackBody = append(animation.AttackBody, rectangle)
			case 0xff0080ff:
				animation.AttackHand = append(animation.AttackHand, rectangle)
			case 0xff00ffff:
				animation.AttackFoot = append(animation.AttackFoot, rectangle)
			default:
				return fmt.Errorf("unrecognize color %x at local point (%d,%d)", rgba, dx, dy)

			}

			rectMin, rectMax := rectangle.Min, rectangle.Max
			for x0 := rectMin.X; x0 <= rectMax.X; x0++ {
				for y0 := rectMin.Y; y0 <= rectMax.Y; y0++ {
					skipped[fmt.Sprintf("%d_%d", x0, y0)] = struct{}{}
				}
			}
		}
	}

	// 角色是否在地面
	animation.IsLand = isLand

	// 粗略计算角色所在的中心位置
	animation.setPosition()

	// JSON
	if err := json.NewEncoder(os.Stdout).Encode(animation); err != nil {
		return err
	}

	return errors.New("TODO :: only Read One File")
}

// 以此点为初始点，向右（X轴）寻找最大坐标maxX，向上（Y轴）寻找最大坐标maxY
func findRectangle(img image.Image, x, y, dx, dy int, rgba uint32) image.Rectangle {
	min, max := image.Point{X: dx, Y: dy}, image.Point{X: dx, Y: dy}

	// X轴寻找
	for x0 := dx; x0 < img.Bounds().Dx(); x0++ {
		r, g, b, a := img.At(x0, y).RGBA()
		if a == 0 {
			break
		}

		rgba0 := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
		if rgba0 != rgba {
			break
		}

		max.X = x0
	}

	// Y轴寻找
	for y0 := dy; y0 < img.Bounds().Dy(); y0++ {
		r, g, b, a := img.At(x, img.Bounds().Dy()-1-y0).RGBA()
		if a == 0 {
			break
		}

		rgba0 := r>>8<<24 | g>>8<<16 | b>>8<<8 | a>>8
		if rgba0 != rgba {
			break
		}

		max.Y = y0
	}

	return image.Rectangle{Min: min, Max: max}
}

func main() {
	if err := filepath.Walk("assets", Walk); err != nil {
		log.Fatalf("filepath.Walk Fail :: %s \n", err.Error())
	}
}
