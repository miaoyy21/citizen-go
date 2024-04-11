package conf

import (
	"fmt"
	"math"
)

type Card struct {
	Id        int       `json:"id"`
	Name      Lang      `json:"name"`
	Level     CardLevel `json:"level"`
	Attribute Attribute `json:"attribute"`
	Value     int       `json:"value"`
	Price     int       `json:"price"`
}

func NewCards(language Language) []Card {
	cards := make([]Card, 0)

	for _, level := range CardLevels {
		for _, attribute := range Attributes {

			// 多语言
			name := NewLang(
				fmt.Sprintf("Lv.%d %s", level, language.Get(LangZhCN, fmt.Sprintf("attribute_%d", attribute))),
				fmt.Sprintf("Lv.%d %s", level, language.Get(LangZhTW, fmt.Sprintf("attribute_%d", attribute))),
				fmt.Sprintf("Lv.%d %s", level, language.Get(LangEnUS, fmt.Sprintf("attribute_%d", attribute))),
			)

			price := CoefficientValues.CardPrice[attribute]

			card := Card{
				Id:        int(level)*1000 + int(attribute),
				Name:      name,
				Level:     level,
				Attribute: attribute,
				Value:     int(CoefficientValues.CardAttributes[attribute] + float64(level-5)*CoefficientValues.CardAttributes[attribute]*CoefficientValues.CardAttributeSteps),
				Price:     int(price * math.Pow(CoefficientValues.CardLevelPrice, float64(level))),
			}

			cards = append(cards, card)
		}
	}

	return cards
}
