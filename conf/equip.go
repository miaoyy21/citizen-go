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

type Equip struct {
	Id         int               `json:"id"`         // 模版ID
	Name       Lang              `json:"name"`       // 装备名称【多语言】
	Saleable   bool              `json:"saleable"`   // 是否可出售
	Level      EquipLevel        `json:"level"`      // 等级
	Attributes map[Attribute]int `json:"attributes"` // 属性
	Price      int               `json:"price"`      // 出售价格

	QualityRate    map[EquipQuality]int    `json:"quality_rate"`     // 不同装备品质的概率 100*10000 = 100%
	NaturalQtyRate map[EquipNaturalQty]int `json:"natural_qty_rate"` // 不同数量天然属性的概率 100*10000 = 100%
}

func NewEquips() []Equip {
	equips := make([]Equip, 0, len(EquipLevels))

	for _, level := range EquipLevels {
		// 基础属性
		attributes := make(map[Attribute]int)
		for attr, value := range CoefficientValues.Attributes {
			attributes[attr] = int(value * CoefficientValues.EquipAttributes[level])
		}

		// 不同品质的概率
		kvQfs := make(map[int]float64)
		for k, v := range CoefficientValues.EquipQuality[level] {
			kvQfs[int(k)] = v
		}

		kvQs, qualityRate := float2int(kvQfs), make(map[EquipQuality]int)
		for k, v := range kvQs {
			qualityRate[EquipQuality(k)] = v
		}

		// 天然属性条数概率
		kvNfs := make(map[int]float64)
		for k, v := range CoefficientValues.EquipNatural[level] {
			kvNfs[int(k)] = v
		}

		kvNs, naturalRate := float2int(kvNfs), make(map[EquipNaturalQty]int)
		for k, v := range kvNs {
			naturalRate[EquipNaturalQty(k)] = v
		}

		equip := Equip{
			Id:             1000 + int(level),
			Name:           NewLang("披风", "披風", "Cloak"),
			Saleable:       true,
			Level:          level,
			Attributes:     attributes,
			Price:          CoefficientValues.EquipPrice[level],
			QualityRate:    qualityRate,
			NaturalQtyRate: naturalRate,
		}

		equips = append(equips, equip)
	}

	return equips
}
