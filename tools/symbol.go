package tools

type Symbol string

var (
	SymbolEnemy Symbol = "enemy"
	SymbolSelf  Symbol = "self"
)

type Direction string

var (
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)

type Effect string

var (
	EffectRed   Effect = "red"
	EffectGreen Effect = "green"
	EffectBlue  Effect = "blue"
)

type SoundCategory string

var (
	SoundCategoryHandStart  SoundCategory = "HandStart"
	SoundCategoryHandEnd    SoundCategory = "HandEnd"
	SoundCategoryFootStart  SoundCategory = "FootStart"
	SoundCategoryFootEnd    SoundCategory = "FootEnd"
	SoundCategorySwingStart SoundCategory = "SwingStart"
	SoundCategorySwingEnd   SoundCategory = "SwingEnd"
)

type SoundCategoryOut string

var (
	SoundCategoryOutHand  SoundCategoryOut = "Hand"
	SoundCategoryOutFoot  SoundCategoryOut = "Foot"
	SoundCategoryOutSwing SoundCategoryOut = "Swing"
)
