package tools

type AttackFrame struct {
	Start int
	End   int

	Powerful  bool
	direction AttackDirection
}

type AttackDirection string

var (
	AttackDirectionUp       = "Up"
	AttackDirectionDown     = "Down"
	AttackDirectionForward  = "Forward"
	AttackDirectionBackward = "Backward"
)

type AttackType string

var (
	AttackTypeInvalid AttackType = "Invalid"
	AttackTypeHead    AttackType = "Head"
	AttackTypeBody    AttackType = "Body"
	AttackTypeHand    AttackType = "Hand"
	AttackTypeFoot    AttackType = "Foot"
)
