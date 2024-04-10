package conf

type EquipLevel int

var (
	EquipLevel1  EquipLevel = 1
	EquipLevel10 EquipLevel = 10
	EquipLevel20 EquipLevel = 20
	EquipLevel30 EquipLevel = 30
	EquipLevel40 EquipLevel = 40
	EquipLevel50 EquipLevel = 50
	EquipLevel60 EquipLevel = 60
)
var EquipLevels = []EquipLevel{
	EquipLevel1, EquipLevel10, EquipLevel20, EquipLevel30,
	EquipLevel40, EquipLevel50, EquipLevel60,
}

type EquipNaturalQty int

var (
	EquipNaturalQty0 EquipNaturalQty = 0
	EquipNaturalQty1 EquipNaturalQty = 1
	EquipNaturalQty2 EquipNaturalQty = 2
	EquipNaturalQty3 EquipNaturalQty = 3
	EquipNaturalQty4 EquipNaturalQty = 4
	EquipNaturalQty5 EquipNaturalQty = 5
)

type EquipQuality int

var (
	EquipQuality0 EquipQuality = 0 // 普通【白】
	EquipQuality1 EquipQuality = 1 // 精良【绿】
	EquipQuality2 EquipQuality = 2 // 卓越【蓝】
	EquipQuality3 EquipQuality = 3 // 史诗【紫】
	EquipQuality4 EquipQuality = 4 // 传说【橙】
	EquipQuality5 EquipQuality = 5 // 神话【红】
)

type EquipColor int

var (
	_                EquipColor = 0 // 无颜色
	EquipColorRed    EquipColor = 1 // 红色
	EquipColorOrange EquipColor = 2 // 橙色
	EquipColorYellow EquipColor = 3 // 黄色
	EquipColorGreen  EquipColor = 4 // 绿色
	EquipColorCyan   EquipColor = 5 // 青色
	EquipColorBlue   EquipColor = 6 // 蓝色
	EquipColorPurple EquipColor = 7 // 紫色
)
