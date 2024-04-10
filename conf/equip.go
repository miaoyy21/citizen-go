package conf

import "fmt"

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
	EquipColorRed    EquipColor = 1 // 红色
	EquipColorOrange EquipColor = 2 // 橙色
	EquipColorYellow EquipColor = 3 // 黄色
	EquipColorGreen  EquipColor = 4 // 绿色
	EquipColorCyan   EquipColor = 5 // 青色
	EquipColorBlue   EquipColor = 6 // 蓝色
	EquipColorPurple EquipColor = 7 // 紫色
)

type Equip struct {
	Id         int               `json:"id"`         // 模版ID
	Name       Lang              `json:"name"`       // 装备名称【多语言】
	Saleable   bool              `json:"saleable"`   // 是否可出售
	Level      EquipLevel        `json:"level"`      // 等级
	Attributes map[Attribute]int `json:"attributes"` // 属性
	Price      int               `json:"price"`      // 出售价格

	ColorRate      map[int]int `json:"color_rate"`       // 装备颜色的概率  100*10000 = 100%
	QualityRate    map[int]int `json:"quality_rate"`     // 不同装备品质的概率 100*10000 = 100%
	NaturalQtyRate map[int]int `json:"natural_qty_rate"` // 不同数量天然属性的概率 100*10000 = 100%
}

func NewEquips() []Equip {
	equips := make([]Equip, 0, len(EquipLevels))

	for _, level := range EquipLevels {
		// 基础属性
		attributes := make(map[Attribute]int)
		for attr, value := range CoefficientValues.Attributes {
			attributes[attr] = int(value * CoefficientValues.EquipAttributes[level])
		}

		// 装备颜色的概率
		kvCfs := make(map[int]float64)
		for k, v := range CoefficientValues.EquipColor {
			kvCfs[int(k)] = v
		}

		// 不同品质的概率
		kvQfs := make(map[int]float64)
		for k, v := range CoefficientValues.EquipQuality[level] {
			kvQfs[int(k)] = v
		}

		// 天然属性条数概率
		kvNfs := make(map[int]float64)
		for k, v := range CoefficientValues.EquipNatural[level] {
			kvNfs[int(k)] = v
		}

		// 多语言
		name := NewLang(
			fmt.Sprintf("Lv.%d %s", level, "披风"),
			fmt.Sprintf("Lv.%d %s", level, "披風"),
			fmt.Sprintf("Lv.%d %s", level, "Cloak"),
		)

		equip := Equip{
			Id:             1000 + int(level),
			Name:           name,
			Saleable:       true,
			Level:          level,
			Attributes:     attributes,
			Price:          CoefficientValues.EquipPrice[level],
			ColorRate:      float2int(kvCfs),
			QualityRate:    float2int(kvQfs),
			NaturalQtyRate: float2int(kvNfs),
		}

		equips = append(equips, equip)
	}

	return equips
}
