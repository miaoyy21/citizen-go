package tools

import (
	"encoding/json"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunSkills(srcAssets, dstAssets string) error {
	// 1. 清空临时文件夹
	if err := clean(filepath.Join(srcAssets, "temp")); err != nil {
		return err
	}
	log.Printf("完成清空文件夹%q ... \n", filepath.Join(srcAssets, "temp"))

	// 2. 清空目标文件夹目录
	for _, d := range []Direction{DirectionLeft, DirectionRight} {
		// 1.1
		if err := clean(filepath.Join(dstAssets, "images", string(d), "enemy", "cape")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "enemy", "cape"))

		// 1.2
		if err := clean(filepath.Join(dstAssets, "images", string(d), "enemy", "effect")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "enemy", "effect"))

		// 1.3
		if err := clean(filepath.Join(dstAssets, "images", string(d), "enemy", "stick")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "enemy", "stick"))

		// 2.1
		if err := clean(filepath.Join(dstAssets, "images", string(d), "self", "cape")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "self", "cape"))

		// 2.2
		if err := clean(filepath.Join(dstAssets, "images", string(d), "self", "effect")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "self", "effect"))

		// 2.3
		if err := clean(filepath.Join(dstAssets, "images", string(d), "self", "stick")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(string(d), "self", "stick"))
	}

	// 3.【碰撞层】：根据角色【只设计向右站位】，产生向左的图片
	if err := filepath.Walk(filepath.Join(srcAssets, "self", "collision"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		return reverse(path, SymbolSelf, filepath.Join(srcAssets, "temp"))
	}); err != nil {
		return err
	}
	log.Printf("完成在文件夹%q，生成角色的左右动画帧解析的样本 ... \n", filepath.Join(srcAssets, "temp"))

	// 4.【碰撞层】：根据敌方单位【只设计向左站位】，产生向左的图片
	if err := filepath.Walk(filepath.Join(srcAssets, "enemy", "collision"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(info.Name(), ".png") {
			return nil
		}

		return reverse(path, SymbolEnemy, filepath.Join(srcAssets, "temp"))
	}); err != nil {
		return err
	}
	log.Printf("完成在文件夹%q，生成敌方单位的左右动画帧解析的样本 ... \n", filepath.Join(srcAssets, "temp"))

	// 5.【碰撞层】：对每个文件进行解析
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
	log.Printf("完成文件夹%q的角色和敌方单位动画帧解析 ... \n", filepath.Join(srcAssets, "temp"))

	// 6.【碰撞层】：生成animations.json文件
	animations := make(map[string]*Animation)
	for _, frame := range frames {
		animation, ok := animations[frame.Name]
		if !ok {
			animation = &Animation{
				Width:  frame.Width,
				Height: frame.Height,

				BreakPrepare:     true,
				LeftSelfFrames:   make([]*Frame, 0),
				RightSelfFrames:  make([]*Frame, 0),
				LeftEnemyFrames:  make([]*Frame, 0),
				RightEnemyFrames: make([]*Frame, 0),

				Files: make(map[string]string),
			}
		}

		if frame.Direction == DirectionLeft && frame.Symbol == SymbolSelf {
			animation.LeftSelfFrames = append(animation.LeftSelfFrames, frame)
			sort.Slice(animation.LeftSelfFrames, func(i, j int) bool {
				return animation.LeftSelfFrames[i].Sequence < animation.LeftSelfFrames[j].Sequence
			})
		} else if frame.Direction == DirectionRight && frame.Symbol == SymbolSelf {
			animation.RightSelfFrames = append(animation.RightSelfFrames, frame)
			sort.Slice(animation.RightSelfFrames, func(i, j int) bool {
				return animation.RightSelfFrames[i].Sequence < animation.RightSelfFrames[j].Sequence
			})
		} else if frame.Direction == DirectionLeft && frame.Symbol == SymbolEnemy {
			animation.LeftEnemyFrames = append(animation.LeftEnemyFrames, frame)
			sort.Slice(animation.LeftEnemyFrames, func(i, j int) bool {
				return animation.LeftEnemyFrames[i].Sequence < animation.LeftEnemyFrames[j].Sequence
			})
		} else {
			animation.RightEnemyFrames = append(animation.RightEnemyFrames, frame)
			sort.Slice(animation.RightEnemyFrames, func(i, j int) bool {
				return animation.RightEnemyFrames[i].Sequence < animation.RightEnemyFrames[j].Sequence
			})
		}

		animations[frame.Name] = animation
	}

	// 6. 7. 计算所处阶段
	for _, animation := range animations {
		animation.LeftSelfFrames = parseStep(animation.LeftSelfFrames)
		animation.RightSelfFrames = parseStep(animation.RightSelfFrames)

		var start, end int
		for _, frame := range animation.RightSelfFrames {
			if frame.Step == StepStart {
				start = frame.Position.X
			}

			if frame.Step == StepHit {
				end = frame.Position.X
				break
			}
		}

		// 是否可以跳过准备阶段
		var y0 int
		for _, frame := range animation.RightSelfFrames {
			if frame.Step != StepPrepare {
				continue
			}

			if y0 == 0 {
				y0 = animation.RightSelfFrames[0].Position.Y
				continue
			}

			if math.Abs(float64(frame.Position.Y-y0)) > 24 {
				animation.BreakPrepare = false
				break
			}
		}

		animation.Distance = end - start
	}

	// 7. 【文件拷贝】
	for _, ds := range [][]string{
		{"self", "cape"},
		{"self", "effect"},
		{"self", "stick"},
		{"enemy", "cape"},
		{"enemy", "effect"},
		{"enemy", "stick"},
	} {
		if err := clean(filepath.Join(srcAssets, "temp")); err != nil {
			return err
		}
		log.Printf("完成清空文件夹%q ... \n", filepath.Join(srcAssets, "temp"))

		src := make([]string, 0)
		src = append(src, srcAssets)
		src = append(src, ds...)
		if err := CopyDirectory(filepath.Join(src...), filepath.Join(srcAssets, "temp")); err != nil {
			return err
		}

		if err := filepath.Walk(filepath.Join(srcAssets, "temp"), func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !strings.HasSuffix(info.Name(), ".png") {
				return nil
			}

			dstLeft := make([]string, 0)
			dstLeft = append(dstLeft, dstAssets)
			dstLeft = append(dstLeft, "images", "left")
			dstLeft = append(dstLeft, ds...)
			if err := onlyReverse(path, filepath.Join(dstLeft...)); err != nil {
				return err
			}

			animation := animations[strings.Split(filepath.Base(path), "_")[0]]

			animation.Files[strings.Join(ds, "_")] = filepath.Join(ds...)
			return nil
		}); err != nil {
			return err
		}

		dstRight := make([]string, 0)
		dstRight = append(dstRight, dstAssets)
		dstRight = append(dstRight, "images", "right")
		dstRight = append(dstRight, ds...)

		if err := CopyDirectory(filepath.Join(src...), filepath.Join(dstRight...)); err != nil {
			return err
		}
		log.Printf("发布动画帧%q至目标目录  ... \n", filepath.Join(ds...))
	}

	// 7.8. 修改释放技能特效的颜色
	for _, ds := range [][]string{
		{"left", "self", "effect"},
		{"right", "self", "effect"},
	} {
		dst := make([]string, 0, len(ds)+2)
		dst = append(dst, dstAssets)
		dst = append(dst, "images")
		dst = append(dst, ds...)

		if err := filepath.Walk(filepath.Join(dst...), func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !strings.HasSuffix(info.Name(), ".png") {
				return nil
			}

			return changeEffect(path, EffectRed)
		}); err != nil {
			return err
		}
		log.Printf("完成在文件夹%q，生成新的特效 ... \n", filepath.Join(dst...))
	}

	// 8.【碰撞层】：拷贝animations.json文件
	file, err := os.Create(filepath.Join(dstAssets, "animations.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(animations); err != nil {
		return err
	}
	log.Printf("发布动画帧解析文件至%q ... \n", filepath.Join(dstAssets, "animations.json"))

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
