package tools

//
//type AttackFrame struct {
//	Start int `json:"start"`
//	End   int `json:"end"`
//
//	AttackType AttackType `json:"type"`
//
//	VSpeed      int `json:"v_speed"`
//	VAccelerate int `json:"v_accelerate"`
//
//	HSpeed      int `json:"h_speed"`
//	HAccelerate int `json:"h_accelerate"`
//}
//
//type AttackType string
//
//var (
//	AttackTypeInvalid AttackType = "Invalid"
//	AttackTypeHead    AttackType = "Head"
//	AttackTypeBody    AttackType = "Body"
//	AttackTypeHand    AttackType = "Hand"
//	AttackTypeFoot    AttackType = "Foot"
//)
//
//func parseAttackFrames(animation *Animation) []*AttackFrame {
//	// 攻击帧信息
//	aFrames := make([]*AttackFrame, 0)
//	attackType, start := AttackTypeInvalid, 0
//
//	// 1. 连续的2个攻击序列中，至少要包含1个非攻击帧
//	for _, frame := range animation.RightFrames {
//		// 新的攻击帧
//		if len(frame.AttackHead) > 0 {
//			if attackType != AttackTypeHead {
//				if attackType != AttackTypeInvalid {
//					aFrames = append(aFrames, &AttackFrame{
//						Start: start,
//						End:   frame.Sequence,
//
//						AttackType: attackType,
//					})
//				}
//
//				attackType = AttackTypeHead
//				start = frame.Sequence
//			}
//		} else if len(frame.AttackBody) > 0 {
//			if attackType != AttackTypeBody {
//				if attackType != AttackTypeInvalid {
//					aFrames = append(aFrames, &AttackFrame{
//						Start: start,
//						End:   frame.Sequence,
//
//						AttackType: attackType,
//					})
//				}
//
//				attackType = AttackTypeBody
//				start = frame.Sequence
//			}
//		} else if len(frame.AttackHand) > 0 {
//			if attackType != AttackTypeHand {
//				if attackType != AttackTypeInvalid {
//					aFrames = append(aFrames, &AttackFrame{
//						Start: start,
//						End:   frame.Sequence,
//
//						AttackType: attackType,
//					})
//				}
//
//				attackType = AttackTypeHand
//				start = frame.Sequence
//			}
//		} else if len(frame.AttackFoot) > 0 {
//			if attackType != AttackTypeFoot {
//				if attackType != AttackTypeInvalid {
//					aFrames = append(aFrames, &AttackFrame{
//						Start: start,
//						End:   frame.Sequence,
//
//						AttackType: attackType,
//					})
//				}
//
//				attackType = AttackTypeFoot
//				start = frame.Sequence
//			}
//		} else if start > 0 {
//			aFrames = append(aFrames, &AttackFrame{
//				Start: start,
//				End:   frame.Sequence - 1,
//
//				AttackType: attackType,
//			})
//
//			attackType = AttackTypeInvalid
//			start = 0
//		}
//	}
//
//	return aFrames
//}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
