package conf

type MateCategory int

var (
	MateCategoryMineralOriginal MateCategory = 11 // 原矿石
	MateCategoryMineralSemi     MateCategory = 12 // 半成品矿石
	MateCategoryMineralFinished MateCategory = 13 // 成品矿石
)
