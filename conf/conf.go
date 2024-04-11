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
	language := NewLanguage()

	return Configuration{
		Language: language,
		Player:   NewPlayer(),
		Equips:   NewEquips(language),
		Cards:    NewCards(language),
		Props:    NewProps(language),
		Mates:    NewMates(language),
	}
}
