package tools

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Frame struct {
	Name   string `json:"-"`
	Width  int    `json:"-"`
	Height int    `json:"-"`

	Sequence int             // 帧顺序号，从1开始
	IsLand   bool            // 是否在地面
	Position image.Point     // 角色所在的粗略中心位置
	Size     image.Rectangle // 真实图片尺寸

	ExposeHead []image.Rectangle // 头（可被他人攻击）
	ExposeBody []image.Rectangle // 体（可被他人攻击）
	ExposeHand []image.Rectangle // 手（可被他人攻击）
	ExposeFoot []image.Rectangle // 脚（可被他人攻击）

	AttackHead []image.Rectangle // 头（可攻击他人）
	AttackBody []image.Rectangle // 体（可攻击他人）
	AttackHand []image.Rectangle // 手（可攻击他人）
	AttackFoot []image.Rectangle // 脚（可攻击他人）
}

func (animation *Frame) rectangle(rss ...[]image.Rectangle) image.Rectangle {
	rects := make([]image.Rectangle, 0)
	for _, rs := range rss {
		rects = append(rects, rs...)
	}

	min, max := image.Point{X: 10000, Y: 10000}, image.Point{}

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

	return image.Rectangle{Min: min, Max: max}
}

func parseFrame(path string) (*Frame, error) {
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
		return nil, fmt.Errorf("[%s] %s", path, "invalid frame file")
	}

	seq, err := strconv.Atoi(n0s[1])
	if err != nil {
		return nil, fmt.Errorf("[%s] %s", path, err.Error())
	}

	//识别矩形框，锚点为左下角
	frame := &Frame{
		Name:   n0s[0],
		Width:  bounds.Dx(),
		Height: bounds.Dy(),

		Sequence:   seq,
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
				frame.ExposeHead = append(frame.ExposeHead, rectangle)
			case 0x00ff00ff:
				frame.ExposeBody = append(frame.ExposeBody, rectangle)
			case 0x000080ff:
				frame.ExposeHand = append(frame.ExposeHand, rectangle)
			case 0x0000ffff:
				frame.ExposeFoot = append(frame.ExposeFoot, rectangle)
			case 0xff8000ff:
				frame.AttackHead = append(frame.AttackHead, rectangle)
			case 0xffff00ff:
				frame.AttackBody = append(frame.AttackBody, rectangle)
			case 0xff0080ff:
				frame.AttackHand = append(frame.AttackHand, rectangle)
			case 0xff00ffff:
				frame.AttackFoot = append(frame.AttackFoot, rectangle)
			default:
				//return nil, fmt.Errorf("[%s] unrecognize color %x at local point (%d,%d)", path, rgba, dx, dy)

			}

			rectMin, rectMax := rectangle.Min, rectangle.Max
			for x0 := rectMin.X; x0 <= rectMax.X; x0++ {
				for y0 := rectMin.Y; y0 <= rectMax.Y; y0++ {
					skipped[fmt.Sprintf("%d_%d", x0, y0)] = struct{}{}
				}
			}
		}
	}

	// 计算真实图片尺寸
	frame.Size = frame.rectangle(
		frame.ExposeHead, frame.ExposeBody, frame.ExposeHand, frame.ExposeFoot,
		frame.AttackHead, frame.AttackBody, frame.AttackHand, frame.AttackFoot,
	)

	// 必须设置角色的身体碰撞框
	if len(frame.ExposeBody)+len(frame.AttackBody) < 1 {
		return nil, fmt.Errorf("[%s] missing Setting Body's Collision", path)
	}

	// 角色是否在地面
	frame.IsLand = isLand

	// 粗略计算角色所在的中心位置
	size := frame.rectangle(frame.ExposeBody, frame.AttackBody)
	frame.Position = image.Point{
		X: int(math.Ceil(float64(size.Min.X+size.Max.X) / 2)),
		Y: int(math.Ceil(float64(size.Min.Y+size.Max.Y) / 2)),
	}

	return frame, nil
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

type Animation struct {
	Width  int
	Height int

	Frames []*Frame
}

func ParseAnimations() error {
	paths := make([]string, 0)
	if err := filepath.Walk("assets", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		if strings.HasPrefix(info.Name(), "stage") {
			return nil
		}

		paths = append(paths, path)
		return nil
	}); err != nil {
		return fmt.Errorf("filepath.parseFrame Fail :: %s ", err.Error())
	}

	frames := make([]*Frame, 0, len(paths))
	for _, path := range paths {
		animation, err := parseFrame(path)
		if err != nil {
			return err
		}

		frames = append(frames, animation)
	}

	// 合并帧为动画
	animations := make(map[string]*Animation)
	for _, frame := range frames {
		animation, ok := animations[frame.Name]
		if !ok {
			animation = &Animation{
				Width:  frame.Width,
				Height: frame.Height,

				Frames: make([]*Frame, 0),
			}
		}

		animation.Frames = append(animation.Frames, frame)
		animations[frame.Name] = animation
	}

	file, err := os.Create("assets/animations.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// JSON
	if err := json.NewEncoder(file).Encode(animations); err != nil {
		return err
	}

	return nil
}
