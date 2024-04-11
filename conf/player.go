package conf

type BirthEquip struct {
	Id         int              `json:"id"`         // 装备模版ID
	Color      EquipColor       `json:"color"`      // 颜色
	Quality    EquipQuality     `json:"quality"`    // 品质
	Attributes []EquipAttribute `json:"attributes"` // 附加属性
}

type BirthCard struct {
	Id  int `json:"id"`  // 卡片模版ID
	Qty int `json:"qty"` // 数量
}

type BirthProp struct {
	Id  int `json:"id"`  // 道具模版ID
	Qty int `json:"qty"` // 数量
}

type BirthMate struct {
	Id  int `json:"id"`  // 材料模版ID
	Qty int `json:"qty"` // 数量
}

// Player 玩家创建时的配置表
type Player struct {
	Gold       int                   `json:"gold"`       // 初始金币
	Attributes map[Attribute]float64 `json:"attributes"` // 初始属性

	Grow   map[Attribute]float64 `json:"grow"`   // 属性成长值，每次升级时对应的属性增长值
	Equips []BirthEquip          `json:"equips"` // 赠送的装备
	Cards  []BirthCard           `json:"cards"`  // 赠送的卡片
	Props  []BirthProp           `json:"props"`  // 赠送的道具
	Mates  []BirthMate           `json:"Mates"`  // 赠送的材料
}

func NewPlayer(conf Configuration) Player {

	// 赠送的装备
	equips := make([]BirthEquip, 0)
	for _, equip := range conf.Equips {
		if equip.birthQty < 1 {
			continue
		}

		for n := 0; n < equip.birthQty; n++ {
			var color EquipColor

			attributes := make([]EquipAttribute, 0)
			if n%3 == 0 {
				color = EquipColorRed
				attributes = append(attributes, EquipAttribute{IsNatural: false, ProtoId: 0})
			} else if n%3 == 1 {
				color = EquipColorGreen
				attributes = append(attributes, EquipAttribute{IsNatural: false, ProtoId: 0}, EquipAttribute{IsNatural: false, ProtoId: 0})
			} else {
				color = EquipColorBlue
				attributes = append(attributes, EquipAttribute{IsNatural: false, ProtoId: 0}, EquipAttribute{IsNatural: false, ProtoId: 0}, EquipAttribute{IsNatural: false, ProtoId: 0})
			}

			equips = append(equips, BirthEquip{
				Id:         equip.Id,
				Color:      color,
				Quality:    EquipQuality2,
				Attributes: attributes,
			})
		}
	}

	// 赠送的卡片
	cards := make([]BirthCard, 0)
	for _, card := range conf.Cards {
		if card.birthQty > 0 {
			cards = append(cards, BirthCard{Id: card.Id, Qty: card.birthQty})
		}
	}

	// 赠送的道具
	props := make([]BirthProp, 0)
	for _, prop := range conf.Props {
		if prop.birthQty > 0 {
			props = append(props, BirthProp{Id: prop.Id, Qty: prop.birthQty})
		}
	}

	// 赠送的材料
	mates := make([]BirthMate, 0)
	for _, mate := range conf.Mates {
		if mate.birthQty > 0 {
			mates = append(mates, BirthMate{Id: mate.Id, Qty: mate.birthQty})
		}
	}

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
		Equips: equips,
		Cards:  cards,
		Props:  props,
		Mates:  mates,
	}
}
