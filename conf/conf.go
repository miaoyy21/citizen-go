package conf

type Configuration struct {
	Language Language `json:"language"`

	Player Player  `json:"player"`
	Equips []Equip `json:"equips"`
	Cards  []Card  `json:"cards"`
	Props  []Prop  `json:"props"`
}

func New() Configuration {
	language := NewLanguage()

	return Configuration{
		Language: language,
		Player:   NewPlayer(),
		Equips:   NewEquips(),
		Cards:    NewCards(language),
		Props:    NewProps(),
	}
}
