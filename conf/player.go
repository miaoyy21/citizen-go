package conf

// Player 玩家创建时的配置表
type Player struct {
	Gold       int                   `json:"gold"`       // 初始金币
	Attributes map[Attribute]float64 `json:"attributes"` // 初始属性
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
	}
}
