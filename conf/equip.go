package conf

type Equip struct {
	Id         int               `json:"id"`         // 模版ID
	Name       Lang              `json:"name"`       // 装备名称【多语言】
	Level      EquipLevel        `json:"level"`      // 等级
	Attributes map[Attribute]int `json:"attributes"` // 属性
	Price      int               `json:"price"`      // 基础售价，实际售价 = 基础售价*装备品质*属性条数

	ColorRate      map[int]int `json:"color_rate"`       // 装备颜色的概率  100*10000 = 100%
	QualityRate    map[int]int `json:"quality_rate"`     // 不同装备品质的概率 100*10000 = 100%
	NaturalQtyRate map[int]int `json:"natural_qty_rate"` // 不同数量天然属性的概率 100*10000 = 100%
}

func NewEquips(language Language) []Equip {
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

		equip := Equip{
			Id:             1000 + int(level),
			Name:           language.Get("equip_cape"),
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
