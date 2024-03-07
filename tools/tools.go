package tools

import (
	"encoding/json"
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
	//attackFrames := make(map[string][]*AttackFrame)
	//for name, animation := range animations {
	//	sizes := make([]image.Rectangle, 0, len(animation.LeftFrames))
	//	for _, frame := range animation.LeftFrames {
	//		sizes = append(sizes, frame.StickSize)
	//	}
	//
	//	animation.Size = rectangle(sizes)
	//
	//	// 合并攻击帧信息
	//	attackFrames[name] = parseAttackFrames(animation)
	//}

	//// 4.5. 仅拷贝没有修改的攻击帧信息
	//for name, aFrames := range attackFrames {
	//	fileName := filepath.Join(dstAssets, "attacks", fmt.Sprintf("%s.json", name))
	//	file, err := os.Open(fileName)
	//	if err != nil {
	//		if os.IsNotExist(err) {
	//			newFile, err := os.Create(fileName)
	//			if err != nil {
	//				return err
	//			}
	//
	//			if err := json.NewEncoder(newFile).Encode(aFrames); err != nil {
	//				newFile.Close()
	//				return err
	//			}
	//
	//			newFile.Close()
	//			continue
	//		}
	//
	//		return err
	//	}
	//
	//	// 读取已发布的攻击帧信息
	//	var oFrames []*AttackFrame
	//	if err := json.NewDecoder(file).Decode(&oFrames); err != nil {
	//		file.Close()
	//		return err
	//	}
	//	file.Close()
	//
	//	// 是否手动维护攻击帧信息
	//	isChange := false
	//	for _, aFrame := range oFrames {
	//		if aFrame.HSpeed != 0 || aFrame.HAccelerate != 0 || aFrame.VSpeed != 0 || aFrame.VAccelerate != 0 {
	//			isChange = true
	//			break
	//		}
	//	}
	//
	//	if isChange {
	//		// 判定攻击帧信息是否变化
	//		sFrame := make([]string, 0, len(aFrames))
	//		for _, aFrame := range aFrames {
	//			sFrame = append(sFrame, fmt.Sprintf("%d_%d_%s", aFrame.Start, aFrame.End, aFrame.AttackType))
	//		}
	//
	//		tFrame := make([]string, 0, len(aFrames))
	//		for _, aFrame := range oFrames {
	//			tFrame = append(tFrame, fmt.Sprintf("%d_%d_%s", aFrame.Start, aFrame.End, aFrame.AttackType))
	//		}
	//
	//		if strings.Join(sFrame, ",") != strings.Join(tFrame, ",") {
	//			log.Printf("[%s] 修改动画帧引起攻击帧信息变更，最新攻击帧信息如下：\n", name)
	//			if err := json.NewEncoder(os.Stdout).Encode(aFrames); err != nil {
	//				return err
	//			}
	//			log.Println()
	//		}
	//	} else {
	//		newFile, err := os.Create(fileName)
	//		if err != nil {
	//			return err
	//		}
	//
	//		if err := json.NewEncoder(newFile).Encode(aFrames); err != nil {
	//			newFile.Close()
	//			return err
	//		}
	//
	//		newFile.Close()
	//	}
	//}

	// 在本目录备份攻击帧信息
	//for name := range attackFrames {
	//	srcFileName := filepath.Join(dstAssets, "attacks", fmt.Sprintf("%s.json", name))
	//	dstFileName := filepath.Join(srcAssets, "attacks", fmt.Sprintf("%s.json", name))
	//	if err := CopyFile(srcFileName, dstFileName); err != nil {
	//		return err
	//	}
	//}

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
