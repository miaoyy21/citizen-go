package conf

type Configuration struct {
	Player Player  `json:"player"`
	Equips []Equip `json:"equips"`
}

func New() Configuration {
	return Configuration{
		Player: NewPlayer(),
		Equips: NewEquips(),
	}
}
