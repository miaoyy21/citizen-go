package conf

type Configuration struct {
	Language Language `json:"language"`

	Player Player  `json:"player"`
	Equips []Equip `json:"equips"`
	Cards  []Card  `json:"cards"`
	Props  []Prop  `json:"props"`
	Mates  []Mate  `json:"mates"`
}

func New() Configuration {
	conf := Configuration{
		Language: NewLanguage(),
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
