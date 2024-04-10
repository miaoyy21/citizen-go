package conf

type Configuration struct {
	Player Player
}

func New() Configuration {
	return Configuration{}
}
