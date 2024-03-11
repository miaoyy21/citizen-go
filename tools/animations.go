package tools

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Frame struct {
	Name string

	Symbol    Symbol
	Direction Direction
	Width     int
	Height    int

	Sequence  int             // 帧顺序号，从1开始
	IsLand    bool            // 是否在地面
	Position  image.Point     // 角色所在的粗略中心位置
	Size      image.Rectangle // 尺寸
	StickSize image.Rectangle // 角色尺寸

	ExposeHead []image.Rectangle // 头（可被他人攻击）
	ExposeBody []image.Rectangle // 体（可被他人攻击）
	ExposeHand []image.Rectangle // 手（可被他人攻击）
	ExposeFoot []image.Rectangle // 脚（可被他人攻击）

	AttackHand []image.Rectangle // 手（可攻击他人）
	AttackFoot []image.Rectangle // 脚（可攻击他人）
}

func rectangle(rss ...[]image.Rectangle) image.Rectangle {
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
	if len(n0s) != 4 {
		return nil, fmt.Errorf("[%s] %s", path, "invalid frame file")
	}

	seq, err := strconv.Atoi(n0s[3])
	if err != nil {
		return nil, fmt.Errorf("[%s] %s", path, err.Error())
	}

	//识别矩形框，锚点为左下角
	frame := &Frame{
		Name:      n0s[2],
		Symbol:    Symbol(n0s[0]),
		Direction: Direction(n0s[1]),
		Width:     bounds.Dx(),
		Height:    bounds.Dy(),

		Sequence:   seq,
		ExposeHead: make([]image.Rectangle, 0),
		ExposeBody: make([]image.Rectangle, 0),
		ExposeHand: make([]image.Rectangle, 0),
		ExposeFoot: make([]image.Rectangle, 0),

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
				if dy <= 10 {
					isLand = true
				}

				continue
			}

			if _, ok := skipped[fmt.Sprintf("%d_%d", dx, dy)]; ok {
				continue
			}

			rect := findRectangle(img, x, y, dx, dy, rgba)
			switch rgba {
			case 0x008000ff:
				frame.ExposeHead = append(frame.ExposeHead, rect)
			case 0x00ff00ff:
				frame.ExposeBody = append(frame.ExposeBody, rect)
			case 0x000080ff:
				frame.ExposeHand = append(frame.ExposeHand, rect)
			case 0x0000ffff:
				frame.ExposeFoot = append(frame.ExposeFoot, rect)
			case 0xff0080ff:
				frame.AttackHand = append(frame.AttackHand, rect)
			case 0xff00ffff:
				frame.AttackFoot = append(frame.AttackFoot, rect)
			default:
				return nil, fmt.Errorf("[%s] unrecognize color %x at local point (%d,%d)", path, rgba, dx, dy)
			}

			rectMin, rectMax := rect.Min, rect.Max
			for x0 := rectMin.X; x0 <= rectMax.X; x0++ {
				for y0 := rectMin.Y; y0 <= rectMax.Y; y0++ {
					skipped[fmt.Sprintf("%d_%d", x0, y0)] = struct{}{}
				}
			}
		}
	}

	// 尺寸
	frame.StickSize = rectangle(
		frame.ExposeHead, frame.ExposeBody, frame.ExposeHand, frame.ExposeFoot,
		frame.AttackHand, frame.AttackFoot,
	)

	// 其他
	frame.Size = rectangle([]image.Rectangle{frame.StickSize})

	// 必须设置角色的身体碰撞框
	if frame.Symbol == SymbolSelf && len(frame.ExposeBody) < 1 {
		return nil, fmt.Errorf("[%s] missing Setting Body's Collision", path)
	}

	// 角色是否在地面
	frame.IsLand = isLand

	// 粗略计算角色所在的中心位置
	size := rectangle(frame.ExposeBody)
	frame.Position = image.Point{
		X: int(math.Ceil(float64(size.Min.X+size.Max.X) / 2)),
		Y: int(math.Ceil(float64(size.Min.Y+size.Max.Y) / 2)),
	}

	// 移除用于调整玩家位置的点
	exposeBody := make([]image.Rectangle, 0)
	for _, rect := range frame.ExposeBody {
		if rect.Max.X-rect.Min.X <= 5 && rect.Max.Y-rect.Min.Y <= 5 {
			continue
		}

		exposeBody = append(exposeBody, rect)
	}
	frame.ExposeBody = exposeBody

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

	LeftSelfFrames  []*Frame
	LeftEnemyFrames []*Frame // 与角色的站位对应

	RightSelfFrames  []*Frame
	RightEnemyFrames []*Frame // 与角色的站位对应

	Files map[string]string
}
