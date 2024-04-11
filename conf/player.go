package conf

// Player 玩家创建时的配置表
type Player struct {
	Gold       int                   `json:"gold"`       // 初始金币
	Attributes map[Attribute]float64 `json:"attributes"` // 初始属性

	Grow map[Attribute]float64 `json:"grow"` // 属性成长值，每次升级时对应的属性增长值
}

func NewPlayer() Player {
	return Player{
		Gold: 100,
		Attributes: map[Attribute]float64{
			AttributeHealth:         100,
			AttributeEnergy:         50,
			AttributeAttack:         10,
			AttributeDefense:        5,
			AttributePenetration:    5,
			AttributeArmor:          5,
			AttributeCritical:       100,
			AttributeResistCritical: 100,
			AttributeAccuracy:       10000,
			AttributeResistAccuracy: 200,
		},
		Grow: map[Attribute]float64{
			AttributeHealth:         20,
			AttributeEnergy:         8,
			AttributeAttack:         3,
			AttributeDefense:        2,
			AttributePenetration:    4,
			AttributeArmor:          2,
			AttributeCritical:       25,
			AttributeResistCritical: 20,
			AttributeAccuracy:       30,
			AttributeResistAccuracy: 10,
		},
	}
}
