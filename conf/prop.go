package conf

import "fmt"

type Prop struct {
	Id          int  `json:"id"`          // 模版ID
	Name        Lang `json:"name"`        // 道具名称【多语言】
	Saleable    bool `json:"saleable"`    // 是否可出售
	Level       int  `json:"level"`       // 使用等级
	Description Lang `json:"description"` // 描述
	Price       int  `json:"price"`       // 售价

	Effect PropEffect `json:"effect"` // 使用效果
	Values []int      `json:"values"` // 相关数值

	birthQty int // 赠送玩家的数量
}

func NewProps(conf Configuration) []Prop {
	props := make([]Prop, 0)

	/****************************** 不可交易的道具 ******************************/
	// （小）生命药水、（中）生命药水、（大）生命药水
	props = append(props, Prop{
		Id: 1011, Saleable: false, Level: 1,
		Effect: PropEffectHealth, Values: []int{100},
		birthQty: 50,
	})
	props = append(props, Prop{
		Id: 1012, Saleable: false, Level: 10,
		Effect: PropEffectHealth, Values: []int{500},
		birthQty: 30,
	})
	props = append(props, Prop{
		Id: 1013, Saleable: false, Level: 20,
		Effect: PropEffectHealth, Values: []int{1000},
		birthQty: 10,
	})

	// （小）精气补充剂、（中）精气补充剂、（大）精气补充剂
	props = append(props, Prop{
		Id: 1021, Saleable: false, Level: 1,
		Effect: PropEffectEnergy, Values: []int{50},
		birthQty: 40,
	})
	props = append(props, Prop{
		Id: 1022, Saleable: false, Level: 10,
		Effect: PropEffectEnergy, Values: []int{250},
		birthQty: 20,
	})
	props = append(props, Prop{
		Id: 1023, Saleable: false, Level: 20,
		Effect: PropEffectEnergy, Values: []int{500},
		birthQty: 10,
	})

	// 金币大礼包
	props = append(props, Prop{
		Id: 1091, Saleable: false, Level: 10,
		Effect: PropEffectGold, Values: []int{5000},
		birthQty: 1,
	})
	props = append(props, Prop{
		Id: 1092, Saleable: false, Level: 30,
		Effect: PropEffectGold, Values: []int{20000},
		birthQty: 1,
	})
	props = append(props, Prop{
		Id: 1093, Saleable: false, Level: 60,
		Effect: PropEffectGold, Values: []int{100000},
		birthQty: 1,
	})

	/******************************  可交易的道具  ******************************/

	// 设置道具名称和道具描述
	newProps := make([]Prop, 0, len(props))
	for _, prop := range props {
		prop.Name = conf.Language.Get(fmt.Sprintf("prop_name_%d", prop.Id))

		args := sliceI2S(prop.Values)
		prop.Description = conf.Language.Get(fmt.Sprintf("prop_description_%d", prop.Id), args...)

		newProps = append(newProps, prop)
	}

	return newProps
}
