package conf

import (
	"fmt"
)

type Prop struct {
	Id          int    `json:"id"`          // 模版ID
	Name        Lang   `json:"name"`        // 道具名称【多语言】
	Saleable    bool   `json:"saleable"`    // 是否可出售
	Level       int    `json:"level"`       // 使用等级
	Description Lang   `json:"description"` // 描述
	Assets      string `json:"assets"`      // 贴图资源ID
	Price       int    `json:"price"`       // 售价

	Effect PropEffect `json:"effect"` // 使用效果
	Values []int      `json:"values"` // 相关数值

	birthQty int // 赠送玩家的数量
}

func NewProps(conf Configuration) []Prop {
	props := make([]Prop, 0)

	/****************************** 不可交易的道具 ******************************/
	// 【不可交易】1.1 ~ 1.3 （小）强力生命药水、（中）强力生命药水、（大）强力生命药水
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
		Effect: PropEffectHealth, Values: []int{1250},
		birthQty: 10,
	})

	// 【不可交易】2.1 ~ 2.3 （小）强力精气补充剂、（中）强力精气补充剂、（大）强力精气补充剂
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
		Effect: PropEffectEnergy, Values: []int{750},
		birthQty: 10,
	})

	// 【不可交易】13.1 ~ 13.3 金币大礼包
	props = append(props, Prop{
		Id: 1131, Saleable: false, Level: 10,
		Effect: PropEffectGold, Values: []int{10000},
		birthQty: 1,
	})
	props = append(props, Prop{
		Id: 1132, Saleable: false, Level: 30,
		Effect: PropEffectGold, Values: []int{50000},
		birthQty: 1,
	})
	props = append(props, Prop{
		Id: 1133, Saleable: false, Level: 60,
		Effect: PropEffectGold, Values: []int{250000},
		birthQty: 1,
	})

	/******************************  可交易的道具  ******************************/
	// 【可交易】1.4 ~ 1.6 （小）生命药水、（中）生命药水、（大）生命药水
	props = append(props, Prop{
		Id: 2014, Saleable: true, Level: 1, Price: 4,
		Effect: PropEffectHealth, Values: []int{50},
	})
	props = append(props, Prop{
		Id: 2015, Saleable: true, Level: 1, Price: 10,
		Effect: PropEffectHealth, Values: []int{150},
	})
	props = append(props, Prop{
		Id: 2016, Saleable: true, Level: 1, Price: 25,
		Effect: PropEffectHealth, Values: []int{375},
	})

	// 【可交易】2.4 ~ 2.6 （小）精气补充剂、（中）精气补充剂、（大）精气补充剂
	props = append(props, Prop{
		Id: 2024, Saleable: true, Level: 1, Price: 3,
		Effect: PropEffectEnergy, Values: []int{25},
	})
	props = append(props, Prop{
		Id: 2025, Saleable: true, Level: 1, Price: 8,
		Effect: PropEffectEnergy, Values: []int{75},
	})
	props = append(props, Prop{
		Id: 2026, Saleable: true, Level: 1, Price: 20,
		Effect: PropEffectEnergy, Values: []int{175},
	})

	// 【可交易】3.1 ~ 3.5 其他提升生命值的道具
	props = append(props, Prop{
		Id: 2031, Saleable: true, Level: 1, Price: 8,
		Effect: PropEffectHealth, Values: []int{80},
	})
	props = append(props, Prop{
		Id: 2032, Saleable: true, Level: 1, Price: 25,
		Effect: PropEffectHealth, Values: []int{200},
	})
	props = append(props, Prop{
		Id: 2033, Saleable: true, Level: 1, Price: 75,
		Effect: PropEffectHealth, Values: []int{450},
	})
	props = append(props, Prop{
		Id: 2034, Saleable: true, Level: 1, Price: 200,
		Effect: PropEffectHealth, Values: []int{1000},
	})
	props = append(props, Prop{
		Id: 2035, Saleable: true, Level: 1, Price: 750,
		Effect: PropEffectHealth, Values: []int{2500},
	})

	// 【可交易】4.1 ~ 4.4 提升生命值百分比的道具
	props = append(props, Prop{
		Id: 2041, Saleable: true, Level: 1, Price: 100,
		Effect: PropEffectHealthPercent, Values: []int{2500},
	})
	props = append(props, Prop{
		Id: 2042, Saleable: true, Level: 1, Price: 300,
		Effect: PropEffectHealthPercent, Values: []int{5000},
	})
	props = append(props, Prop{
		Id: 2043, Saleable: true, Level: 1, Price: 750,
		Effect: PropEffectHealthPercent, Values: []int{7500},
	})
	props = append(props, Prop{
		Id: 2044, Saleable: true, Level: 1, Price: 2000,
		Effect: PropEffectHealthPercent, Values: []int{10000},
	})

	// 【可交易】5.1 ~ 5.5 其他提升精气值的道具
	props = append(props, Prop{
		Id: 2051, Saleable: true, Level: 1, Price: 5,
		Effect: PropEffectEnergy, Values: []int{40},
	})
	props = append(props, Prop{
		Id: 2052, Saleable: true, Level: 1, Price: 15,
		Effect: PropEffectEnergy, Values: []int{100},
	})
	props = append(props, Prop{
		Id: 2053, Saleable: true, Level: 1, Price: 30,
		Effect: PropEffectEnergy, Values: []int{225},
	})
	props = append(props, Prop{
		Id: 2054, Saleable: true, Level: 1, Price: 100,
		Effect: PropEffectEnergy, Values: []int{500},
	})
	props = append(props, Prop{
		Id: 2055, Saleable: true, Level: 1, Price: 400,
		Effect: PropEffectEnergy, Values: []int{1250},
	})

	// 【可交易】6.1 ~ 6.4 提升精气值百分比的道具
	props = append(props, Prop{
		Id: 2061, Saleable: true, Level: 1, Price: 60,
		Effect: PropEffectEnergyPercent, Values: []int{2500},
	})
	props = append(props, Prop{
		Id: 2062, Saleable: true, Level: 1, Price: 200,
		Effect: PropEffectEnergyPercent, Values: []int{5000},
	})
	props = append(props, Prop{
		Id: 2063, Saleable: true, Level: 1, Price: 500,
		Effect: PropEffectEnergyPercent, Values: []int{7500},
	})
	props = append(props, Prop{
		Id: 2064, Saleable: true, Level: 1, Price: 1200,
		Effect: PropEffectEnergyPercent, Values: []int{10000},
	})

	// 【可交易】7.1 ~ 7.4 提升攻击力的道具
	props = append(props, Prop{
		Id: 2071, Saleable: true, Level: 1, Price: 75,
		Effect: PropEffectAttack, Values: []int{50, 60},
	})
	props = append(props, Prop{
		Id: 2072, Saleable: true, Level: 1, Price: 225,
		Effect: PropEffectAttack, Values: []int{80, 60},
	})
	props = append(props, Prop{
		Id: 2073, Saleable: true, Level: 1, Price: 600,
		Effect: PropEffectAttack, Values: []int{150, 60},
	})
	props = append(props, Prop{
		Id: 2074, Saleable: true, Level: 1, Price: 1500,
		Effect: PropEffectAttack, Values: []int{250, 60},
	})

	// 【可交易】8.1 ~ 8.4 提升攻击力百分比的道具
	props = append(props, Prop{
		Id: 2081, Saleable: true, Level: 1, Price: 200,
		Effect: PropEffectAttackPercent, Values: []int{5000, 90},
	})
	props = append(props, Prop{
		Id: 2082, Saleable: true, Level: 1, Price: 800,
		Effect: PropEffectAttackPercent, Values: []int{10000, 90},
	})
	props = append(props, Prop{
		Id: 2083, Saleable: true, Level: 1, Price: 3500,
		Effect: PropEffectAttackPercent, Values: []int{20000, 90},
	})
	props = append(props, Prop{
		Id: 2084, Saleable: true, Level: 1, Price: 15000,
		Effect: PropEffectAttackPercent, Values: []int{30000, 90},
	})

	// 【可交易】9.1 ~ 9.4 提升吸收伤害百分比的道具
	props = append(props, Prop{
		Id: 2091, Saleable: true, Level: 1, Price: 150,
		Effect: PropEffectArmorPercent, Values: []int{2500, 30},
	})
	props = append(props, Prop{
		Id: 2092, Saleable: true, Level: 1, Price: 500,
		Effect: PropEffectArmorPercent, Values: []int{5000, 30},
	})
	props = append(props, Prop{
		Id: 2093, Saleable: true, Level: 1, Price: 2500,
		Effect: PropEffectArmorPercent, Values: []int{7500, 30},
	})
	props = append(props, Prop{
		Id: 2094, Saleable: true, Level: 1, Price: 10000,
		Effect: PropEffectArmorPercent, Values: []int{10000, 30},
	})

	// 【可交易】10.1 ~ 10.4 提升暴击百分比的道具
	props = append(props, Prop{
		Id: 2101, Saleable: true, Level: 1, Price: 200,
		Effect: PropEffectCriticalPercent, Values: []int{2500, 45},
	})
	props = append(props, Prop{
		Id: 2102, Saleable: true, Level: 1, Price: 800,
		Effect: PropEffectCriticalPercent, Values: []int{5000, 45},
	})
	props = append(props, Prop{
		Id: 2103, Saleable: true, Level: 1, Price: 3500,
		Effect: PropEffectCriticalPercent, Values: []int{7500, 45},
	})
	props = append(props, Prop{
		Id: 2104, Saleable: true, Level: 1, Price: 15000,
		Effect: PropEffectCriticalPercent, Values: []int{10000, 45},
	})

	// 【可交易】11.1 ~ 11.4 提升暴击百分比的道具
	props = append(props, Prop{
		Id: 2111, Saleable: true, Level: 1, Price: 100,
		Effect: PropEffectAccuracyPercent, Values: []int{2500, 75},
	})
	props = append(props, Prop{
		Id: 2112, Saleable: true, Level: 1, Price: 300,
		Effect: PropEffectAccuracyPercent, Values: []int{5000, 75},
	})
	props = append(props, Prop{
		Id: 2113, Saleable: true, Level: 1, Price: 1000,
		Effect: PropEffectAccuracyPercent, Values: []int{7500, 75},
	})
	props = append(props, Prop{
		Id: 2114, Saleable: true, Level: 1, Price: 5000,
		Effect: PropEffectAccuracyPercent, Values: []int{10000, 75},
	})

	// 【可交易】12.1 可修改角色名称的道具
	props = append(props, Prop{
		Id: 2121, Saleable: true, Level: 1, Price: 25000,
		Effect: PropEffectPlayerName, Values: []int{},
	})

	// 【可交易】12.2 可修改角色颜色的道具
	props = append(props, Prop{
		Id: 2122, Saleable: true, Level: 1, Price: 20000,
		Effect: PropEffectPlayerColor, Values: []int{},
	})

	// 设置道具名称和道具描述
	newProps := make([]Prop, 0, len(props))
	for _, prop := range props {

		// 贴图资源ID
		row, col := prop.Id%1000/10, prop.Id%10
		prop.Assets = fmt.Sprintf("%d_%d", row, col)

		prop.Name = conf.Language.Get(fmt.Sprintf("prop_name_%d", prop.Id))

		args := sliceI2S(prop.Values)
		if prop.Effect == PropEffectAttackPercent || prop.Effect == PropEffectArmorPercent ||
			prop.Effect == PropEffectCriticalPercent || prop.Effect == PropEffectAccuracyPercent {
			args = sliceP2S(prop.Values)
		}

		prop.Description = conf.Language.Get(fmt.Sprintf("prop_description_%d", prop.Id), args...)

		newProps = append(newProps, prop)
	}

	return newProps
}
