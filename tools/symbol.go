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
	SoundCategoryBlowStart  SoundCategory = "BlowStart"
	SoundCategoryBlowEnd    SoundCategory = "BlowEnd"
	SoundCategorySwingStart SoundCategory = "SwingStart"
	SoundCategorySwingEnd   SoundCategory = "SwingEnd"
)

type SoundCategoryOut string

var (
	SoundCategoryOutBlow  SoundCategoryOut = "Blow"
	SoundCategoryOutSwing SoundCategoryOut = "Swing"
)
