package conf

import (
	"fmt"
	"math"
)

type Card struct {
	Id    int       `json:"id"`    // 模版ID
	Name  Lang      `json:"name"`  // 卡片名称【多语言】
	Level CardLevel `json:"level"` // 等级
	Price int       `json:"price"` // 售价

	Attribute Attribute `json:"attribute"` // 属性
	Value     int       `json:"value"`     // 属性值
}

func NewCards(language Language) []Card {
	cards := make([]Card, 0)

	for _, level := range CardLevels {
		for _, attribute := range Attributes {
			price := CoefficientValues.CardPrice[attribute]

			card := Card{
				Id:        int(level)*1000 + int(attribute),
				Name:      language.Get(fmt.Sprintf("attribute_%d", attribute)),
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
