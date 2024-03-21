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
