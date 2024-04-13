package proto

type Proto struct {
	Language    Language              `json:"language"`
	EquipAssets map[EquipColor]string `json:"equip_assets"`

	Player Player  `json:"player"`
	Equips []Equip `json:"equips"`
	Cards  []Card  `json:"cards"`
	Props  []Prop  `json:"props"`
	Mates  []Mate  `json:"mates"`
}

func New() Proto {
	conf := Proto{
		Language: NewLanguage(),
		EquipAssets: map[EquipColor]string{
			EquipColorRed:    "1_1",
			EquipColorOrange: "1_2",
			EquipColorYellow: "1_3",
			EquipColorGreen:  "1_4",
			EquipColorCyan:   "1_5",
			EquipColorBlue:   "1_6",
			EquipColorPurple: "1_7",
		},
	}

	// 装备、卡片、道具、材料
	conf.Equips = NewEquips(conf)
	conf.Cards = NewCards(conf)
	conf.Props = NewProps(conf)
	conf.Mates = NewMates(conf)

	// 玩家
	conf.Player = NewPlayer(conf)

	return conf
}
