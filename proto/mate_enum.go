package proto

type MateCategory int

var (
	MateCategoryMineralOriginal MateCategory = 1 // 原矿石
	MateCategoryMineralSemi     MateCategory = 2 // 半成品矿石
	MateCategoryMineralFinished MateCategory = 3 // 成品矿石
)
