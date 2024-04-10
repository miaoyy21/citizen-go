package conf

// Player 玩家创建时的配置表
type Player struct {
	Gold       int                   // 初始金币
	Attributes map[Attribute]float64 // 初始属性
}

func NewPlayer() Player {
	return Player{
		Gold: 100,
		Attributes: map[Attribute]float64{
			Health:         100,
			Energy:         50,
			Attack:         10,
			Defense:        5,
			Penetration:    5,
			Armor:          5,
			Critical:       100,
			ResistCritical: 100,
			Accuracy:       10000,
			ResistAccuracy: 200,
		},
	}
}
