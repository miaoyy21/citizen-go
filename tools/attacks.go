package tools

type AttackFrame struct {
	Start int `json:"start"`
	End   int `json:"end"`

	AttackType AttackType `json:"type"`

	VSpeed      int `json:"v_speed"`
	VAccelerate int `json:"v_accelerate"`

	HSpeed      int `json:"h_speed"`
	HAccelerate int `json:"h_accelerate"`
}

type AttackType string

var (
	AttackTypeInvalid AttackType = "Invalid"
	AttackTypeHead    AttackType = "Head"
	AttackTypeBody    AttackType = "Body"
	AttackTypeHand    AttackType = "Hand"
	AttackTypeFoot    AttackType = "Foot"
)

func parseAttackFrames(animation *Animation) []*AttackFrame {
	// 攻击帧信息
	aFrames := make([]*AttackFrame, 0)
	attackType, start := AttackTypeInvalid, 0

	// 1. 连续的2个攻击序列中，至少要包含1个非攻击帧
	for _, frame := range animation.RightFrames {
		// 新的攻击帧
		if len(frame.AttackHead) > 0 {
			if attackType != AttackTypeHead {
				if attackType != AttackTypeInvalid {
					aFrames = append(aFrames, &AttackFrame{
						Start: start,
						End:   frame.Sequence,

						AttackType: attackType,
					})
				}

				attackType = AttackTypeHead
				start = frame.Sequence
			}
		} else if len(frame.AttackBody) > 0 {
			if attackType != AttackTypeBody {
				if attackType != AttackTypeInvalid {
					aFrames = append(aFrames, &AttackFrame{
						Start: start,
						End:   frame.Sequence,

						AttackType: attackType,
					})
				}

				attackType = AttackTypeBody
				start = frame.Sequence
			}
		} else if len(frame.AttackHand) > 0 {
			if attackType != AttackTypeHand {
				if attackType != AttackTypeInvalid {
					aFrames = append(aFrames, &AttackFrame{
						Start: start,
						End:   frame.Sequence,

						AttackType: attackType,
					})
				}

				attackType = AttackTypeHand
				start = frame.Sequence
			}
		} else if len(frame.AttackFoot) > 0 {
			if attackType != AttackTypeFoot {
				if attackType != AttackTypeInvalid {
					aFrames = append(aFrames, &AttackFrame{
						Start: start,
						End:   frame.Sequence,

						AttackType: attackType,
					})
				}

				attackType = AttackTypeFoot
				start = frame.Sequence
			}
		} else if start > 0 {
			aFrames = append(aFrames, &AttackFrame{
				Start: start,
				End:   frame.Sequence - 1,

				AttackType: attackType,
			})

			attackType = AttackTypeInvalid
			start = 0
		}
	}

	return aFrames
}
