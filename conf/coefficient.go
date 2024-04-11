package conf

import (
	"fmt"
	"math"
)

type Coefficient struct {
	Attributes         map[Attribute]float64                      // 每种属性的基数
	EquipAttributes    map[EquipLevel]float64                     // 装备等级对应的属性系数
	EquipPrice         map[EquipLevel]int                         // 装备等级对应的基础售价
	EquipColor         map[EquipColor]float64                     // 装备的颜色系数
	EquipQuality       map[EquipLevel]map[EquipQuality]float64    // 装备等级对应的品质系数
	EquipNatural       map[EquipLevel]map[EquipNaturalQty]float64 // 装备等级对应的天然属性系数
	CardAttributes     map[Attribute]float64                      // 卡片的属性，以Lv.5的属性卡为基准
	CardAttributeSteps float64                                    // 卡片的属性衰减数值 公式：b + (x-5)*v*b
	CardPrice          map[Attribute]float64                      // 每种属性卡的Lv.1售价
	CardLevelPrice     float64                                    // 卡片每提升1级，增加的售价系数
}

var CoefficientValues = Coefficient{
	Attributes: map[Attribute]float64{
		AttributeHealth:         25,
		AttributeEnergy:         5,
		AttributeAttack:         4,
		AttributeDefense:        2,
		AttributePenetration:    3,
		AttributeArmor:          2,
		AttributeCritical:       50,
		AttributeResistCritical: 25,
		AttributeAccuracy:       50,
		AttributeResistAccuracy: 75,
	},
	EquipAttributes: map[EquipLevel]float64{
		EquipLevel1:  1.5 * math.Pow(1.25, 0),
		EquipLevel10: 1.5 * math.Pow(1.25, 1),
		EquipLevel20: 1.6 * math.Pow(1.30, 2),
		EquipLevel30: 1.7 * math.Pow(1.35, 3),
		EquipLevel40: 1.8 * math.Pow(1.40, 4),
		EquipLevel50: 1.9 * math.Pow(1.45, 5),
		EquipLevel60: 2.0 * math.Pow(1.50, 6),
	},
	EquipPrice: map[EquipLevel]int{
		EquipLevel1:  int(10 * math.Pow(1.25, 0)),
		EquipLevel10: int(10 * math.Pow(1.25, 1)),
		EquipLevel20: int(11 * math.Pow(1.30, 2)),
		EquipLevel30: int(12 * math.Pow(1.35, 3)),
		EquipLevel40: int(13 * math.Pow(1.40, 4)),
		EquipLevel50: int(14 * math.Pow(1.45, 5)),
		EquipLevel60: int(15 * math.Pow(1.50, 6)),
	},
	EquipColor: map[EquipColor]float64{
		EquipColorRed:    8,
		EquipColorOrange: 1,
		EquipColorYellow: 7,
		EquipColorGreen:  9,
		EquipColorCyan:   6,
		EquipColorBlue:   10,
		EquipColorPurple: 3,
	},
	EquipQuality: map[EquipLevel]map[EquipQuality]float64{
		EquipLevel1: {
			EquipQuality0: math.Pow(0.50, 1),
			EquipQuality1: math.Pow(0.58, 2),
			EquipQuality2: math.Pow(0.66, 3),
			EquipQuality3: math.Pow(0.74, 4),
			EquipQuality4: math.Pow(0.76, 5),
			EquipQuality5: math.Pow(0.78, 6),
		},
		EquipLevel10: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.56, 1),
			EquipQuality2: math.Pow(0.62, 2),
			EquipQuality3: math.Pow(0.68, 3),
			EquipQuality4: math.Pow(0.74, 4),
			EquipQuality5: math.Pow(0.76, 5),
		},
		EquipLevel20: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.55, 1),
			EquipQuality2: math.Pow(0.60, 2),
			EquipQuality3: math.Pow(0.65, 3),
			EquipQuality4: math.Pow(0.70, 4),
			EquipQuality5: math.Pow(0.70, 5),
		},
		EquipLevel30: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.54, 1),
			EquipQuality2: math.Pow(0.58, 2),
			EquipQuality3: math.Pow(0.62, 3),
			EquipQuality4: math.Pow(0.66, 4),
			EquipQuality5: math.Pow(0.62, 5),
		},
		EquipLevel40: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.53, 1),
			EquipQuality2: math.Pow(0.56, 2),
			EquipQuality3: math.Pow(0.59, 3),
			EquipQuality4: math.Pow(0.62, 4),
			EquipQuality5: math.Pow(0.56, 5),
		},
		EquipLevel50: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.52, 1),
			EquipQuality2: math.Pow(0.54, 2),
			EquipQuality3: math.Pow(0.56, 3),
			EquipQuality4: math.Pow(0.58, 4),
			EquipQuality5: math.Pow(0.52, 5),
		},
		EquipLevel60: {
			EquipQuality0: math.Pow(0.50, 0),
			EquipQuality1: math.Pow(0.51, 1),
			EquipQuality2: math.Pow(0.52, 2),
			EquipQuality3: math.Pow(0.53, 3),
			EquipQuality4: math.Pow(0.54, 4),
			EquipQuality5: math.Pow(0.50, 5),
		},
	},
	EquipNatural: map[EquipLevel]map[EquipNaturalQty]float64{
		EquipLevel1: {
			EquipNaturalQty0: 20 * math.Pow(0.5, 0),
			EquipNaturalQty1: 21 * math.Pow(0.51, 1),
			EquipNaturalQty2: 22 * math.Pow(0.52, 2),
			EquipNaturalQty3: 23 * math.Pow(0.53, 3),
			EquipNaturalQty4: 24 * math.Pow(0.54, 4),
			EquipNaturalQty5: 25 * math.Pow(0.55, 5),
		},
		EquipLevel10: {
			EquipNaturalQty0: 20 * math.Pow(0.50, 0),
			EquipNaturalQty1: 19 * math.Pow(0.50, 1),
			EquipNaturalQty2: 18 * math.Pow(0.50, 2),
			EquipNaturalQty3: 17 * math.Pow(0.50, 3),
			EquipNaturalQty4: 16 * math.Pow(0.50, 4),
			EquipNaturalQty5: 15 * math.Pow(0.50, 5),
		},
		EquipLevel20: {
			EquipNaturalQty0: 20 * math.Pow(0.49, 0),
			EquipNaturalQty1: 19 * math.Pow(0.49, 1),
			EquipNaturalQty2: 18 * math.Pow(0.48, 2),
			EquipNaturalQty3: 17 * math.Pow(0.47, 3),
			EquipNaturalQty4: 16 * math.Pow(0.46, 4),
			EquipNaturalQty5: 15 * math.Pow(0.45, 5),
		},
		EquipLevel30: {
			EquipNaturalQty0: 20 * math.Pow(0.47, 0),
			EquipNaturalQty1: 19 * math.Pow(0.47, 1),
			EquipNaturalQty2: 18 * math.Pow(0.45, 2),
			EquipNaturalQty3: 17 * math.Pow(0.43, 3),
			EquipNaturalQty4: 16 * math.Pow(0.41, 4),
			EquipNaturalQty5: 15 * math.Pow(0.39, 5),
		},
		EquipLevel40: {
			EquipNaturalQty0: 20 * math.Pow(0.47, 0),
			EquipNaturalQty1: 18 * math.Pow(0.47, 1),
			EquipNaturalQty2: 16 * math.Pow(0.45, 2),
			EquipNaturalQty3: 14 * math.Pow(0.43, 3),
			EquipNaturalQty4: 12 * math.Pow(0.41, 4),
			EquipNaturalQty5: 10 * math.Pow(0.39, 5),
		},
		EquipLevel50: {
			EquipNaturalQty0: 20 * math.Pow(0.44, 0),
			EquipNaturalQty1: 18 * math.Pow(0.44, 1),
			EquipNaturalQty2: 16 * math.Pow(0.41, 2),
			EquipNaturalQty3: 14 * math.Pow(0.38, 3),
			EquipNaturalQty4: 12 * math.Pow(0.35, 4),
			EquipNaturalQty5: 10 * math.Pow(0.32, 5),
		},
		EquipLevel60: {
			EquipNaturalQty0: 20 * math.Pow(0.44, 0),
			EquipNaturalQty1: 17 * math.Pow(0.44, 1),
			EquipNaturalQty2: 14 * math.Pow(0.41, 2),
			EquipNaturalQty3: 11 * math.Pow(0.38, 3),
			EquipNaturalQty4: 7 * math.Pow(0.35, 4),
			EquipNaturalQty5: 3 * math.Pow(0.32, 5),
		},
	},
	CardAttributes: map[Attribute]float64{
		AttributeHealth:         200,
		AttributeEnergy:         50,
		AttributeAttack:         40,
		AttributeDefense:        20,
		AttributePenetration:    20,
		AttributeArmor:          10,
		AttributeCritical:       750,
		AttributeResistCritical: 350,
		AttributeAccuracy:       1000,
		AttributeResistAccuracy: 500,
	},
	CardAttributeSteps: 0.2,
	CardPrice: map[Attribute]float64{
		AttributeHealth:         16,
		AttributeEnergy:         9,
		AttributeAttack:         30,
		AttributeDefense:        16,
		AttributePenetration:    25,
		AttributeArmor:          12,
		AttributeCritical:       18,
		AttributeResistCritical: 10,
		AttributeAccuracy:       12,
		AttributeResistAccuracy: 8,
	},
	CardLevelPrice: 1.50,
}

func float2int(values map[int]float64) map[int]int {
	var sum float64
	var minKey = 999999

	for key, value := range values {
		if key < minKey {
			minKey = key
		}

		sum = sum + value
	}

	newValues, newSum := make(map[int]int), 0
	for key, value := range values {
		if key == minKey {
			continue
		}

		newValue := int(math.Floor(value * 10000.0 / sum))
		newValues[key] = newValue
		newSum = newSum + newValue
	}

	newValues[minKey] = 10000 - newSum
	return newValues
}

func sliceI2S(values []int) []string {
	ss := make([]string, 0, len(values))

	for _, value := range values {
		ss = append(ss, fmt.Sprintf("%d", value))
	}

	return ss
}
