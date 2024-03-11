package tools

import "image"

type Step string

var (
	StepStart   Step = "start"
	StepPrepare Step = "prepare"
	StepHit     Step = "hit"
	StepFinish  Step = "finish"
)

func parseStep(frames []*Frame) []*Frame {
	newFrames := make([]*Frame, 0, len(frames))

	step := 0
	startPosition := image.Point{}
	for index, frame := range frames {
		if step == 0 {
			step = 1
			startPosition = frame.Position
		}

		if step == 1 {
			if frame.Position.Eq(startPosition) {
				frame.Step = StepStart
			} else {
				step = 2
				frame.Step = StepPrepare
			}
		} else if step == 2 {
			if len(frame.AttackHand) > 0 || len(frame.AttackFoot) > 0 {
				step = 3
				frame.Step = StepHit
			} else {
				frame.Step = StepPrepare
			}
		} else if step == 3 {
			isHit := false
			for _, fNext := range frames[index:] {
				if len(fNext.AttackHand) > 0 || len(fNext.AttackFoot) > 0 {
					isHit = true
					break
				}
			}

			if !isHit {
				step = 4
				frame.Step = StepFinish
			} else {
				frame.Step = StepHit
			}
		} else {
			frame.Step = StepFinish
		}

		newFrames = append(newFrames, frame)
	}

	return newFrames
}
