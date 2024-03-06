package tools

import (
	"encoding/json"
	"image"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Run(srcAssets, dstAssets string) error {
	// 1.【碰撞层】：清空临时文件夹
	if err := clean(filepath.Join(srcAssets, "temp")); err != nil {
		return err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(srcAssets, "temp"))
	if err := clean(filepath.Join(dstAssets, "images")); err != nil {
		return err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(dstAssets, "images"))

	// 2.【碰撞层】：根据角色站位【只设计向右站位】，产生向右的图片
	if err := filepath.Walk(filepath.Join(srcAssets, "collision"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		if err := reverse(path, SymbolCollision, filepath.Join(srcAssets, "temp")); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	log.Printf("完成在文件夹%q，生成左右动画帧解析的样本 ... \n", filepath.Join(srcAssets, "temp"))

	// 3.【碰撞层】：对每个文件进行解析
	frames := make([]*Frame, 0)
	if err := filepath.Walk(filepath.Join(srcAssets, "temp"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		frame, err := parseFrame(path)
		if err != nil {
			return err
		}

		frames = append(frames, frame)

		return nil
	}); err != nil {
		return err
	}
	log.Printf("完成文件夹%q的所有动画帧解析 ... \n", filepath.Join(srcAssets, "temp"))

	// 4.【碰撞层】：生成animations.json文件
	animations := make(map[string]*Animation)
	for _, frame := range frames {
		animation, ok := animations[frame.Name]
		if !ok {
			animation = &Animation{
				Width:  frame.Width,
				Height: frame.Height,

				LeftFrames:  make([]*Frame, 0),
				RightFrames: make([]*Frame, 0),
			}
		}

		if frame.Direction == DirectionLeft {
			animation.LeftFrames = append(animation.LeftFrames, frame)
			sort.Slice(animation.LeftFrames, func(i, j int) bool {
				return animation.LeftFrames[i].Sequence < animation.LeftFrames[j].Sequence
			})
		} else {
			animation.RightFrames = append(animation.RightFrames, frame)
			sort.Slice(animation.RightFrames, func(i, j int) bool {
				return animation.RightFrames[i].Sequence < animation.RightFrames[j].Sequence
			})
		}

		animations[frame.Name] = animation
	}

	// 计算整个动画尺寸
	for _, animation := range animations {
		sizes := make([]image.Rectangle, 0, len(animation.LeftFrames))
		for _, frame := range animation.LeftFrames {
			sizes = append(sizes, frame.StickSize)
		}

		animation.Size = rectangle(sizes)

		//// 攻击帧信息
		//attackFrames := make([]*AttackFrame, 0)
		//attackType, start := AttackTypeInvalid, 0
		//startPos, attackPos := image.Point{}, image.Point{}
		//
		//// 1. 连续的2个攻击序列中，至少要包含1个非攻击帧
		//// 2. 一个攻击序列内，如果帧数大于等于4个，那么认为是强力攻击
		//// 3. 根据角色的中心初始（攻击开始帧的前1帧）与结束位置判定攻击的垂直方向是向上还是向下
		//// 4. 根据角色的攻击开始与结束位置
		//// 5. 如果攻击帧仅有1帧，那么通过攻击位置和中心位置进行判定攻击方向
		//for _, frame := range animation.RightFrames {
		//	if len(frame.AttackHead) > 0 {
		//		attackRect := rectangle(frame.AttackHead)
		//		attackPos = image.Point{
		//			X: int(math.Ceil(float64(attackRect.Min.X+attackRect.Max.X) / 2)),
		//			Y: int(math.Ceil(float64(attackRect.Min.Y+attackRect.Max.Y) / 2)),
		//		}
		//
		//		if attackType != AttackTypeHead {
		//			// 新的攻击帧
		//			attackType = AttackTypeHead
		//			start = frame.Sequence
		//			startPos = frame.Position
		//
		//		} else {
		//			// 攻击帧位置偏移量来判定是否属于新的攻击帧
		//			rectangle(frame.AttackHead)
		//		}
		//	}
		//	if len(frame.AttackBody) > 0 {
		//		attackRect := rectangle(frame.AttackHead)
		//		attackPos = image.Point{
		//			X: int(math.Ceil(float64(attackRect.Min.X+attackRect.Max.X) / 2)),
		//			Y: int(math.Ceil(float64(attackRect.Min.Y+attackRect.Max.Y) / 2)),
		//		}
		//
		//		if attackType != AttackTypeHead {
		//			// 新的攻击帧
		//			attackType = AttackTypeHead
		//			start = frame.Sequence
		//			startPos = frame.Position
		//
		//		} else {
		//			// 攻击帧位置偏移量来判定是否属于新的攻击帧
		//			rectangle(frame.AttackHead)
		//		}
		//	} else if start > 0 {
		//		attackFrames = append(attackFrames, &AttackFrame{
		//			Start: start,
		//			End:   frame.Sequence - 1,
		//		})
		//	}
		//}
	}

	// 5.【碰撞层】：拷贝animations.json文件
	file, err := os.Create(filepath.Join(srcAssets, "animations.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(animations); err != nil {
		return err
	}

	if err := CopyFile(filepath.Join(srcAssets, "animations.json"), filepath.Join(dstAssets, "animations.json")); err != nil {
		return err
	}
	log.Printf("发布动画帧解析文件至%q ... \n", filepath.Join(dstAssets, "animations.json"))

	// 6.【角色层】：清空临时文件夹
	if err := clean(filepath.Join(srcAssets, "temp")); err != nil {
		return err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(srcAssets, "temp"))

	// 7.【碰撞层】：根据角色站位【只设计向右站位】，产生向右的图片
	if err := filepath.Walk(filepath.Join(srcAssets, "stick"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		if err := reverse(path, SymbolCollision, filepath.Join(srcAssets, "temp")); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	log.Printf("完成在文件夹%q，生成左右动画帧 ... \n", filepath.Join(srcAssets, "temp"))

	// 8.拷贝动画帧到指定目录
	if err := CopyDirectory(filepath.Join(srcAssets, "temp"), filepath.Join(dstAssets, "images")); err != nil {
		return err
	}
	log.Printf("发布所有动画帧至目录%q ... \n", filepath.Join(dstAssets, "images"))

	// 9. 拷贝empty.png文件
	if err := CopyFile(filepath.Join(srcAssets, "empty.png"), filepath.Join(dstAssets, "images", "empty.png")); err != nil {
		return err
	}
	log.Printf("发布空图文件至%q ... \n", filepath.Join(dstAssets, "images", "empty.png"))

	// 10. 拷贝地图文件
	if err := CopyDirectory(filepath.Join(srcAssets, "stage"), filepath.Join(dstAssets, "images")); err != nil {
		return err
	}
	log.Printf("发布所有地图文件至目录%q ... \n", filepath.Join(dstAssets, "images"))

	return nil
}
