package conf

import (
	"fmt"
	"math"
)

type Card struct {
	Id     int       `json:"id"`     // 模版ID
	Name   Lang      `json:"name"`   // 卡片名称【多语言】
	Level  CardLevel `json:"level"`  // 等级
	Assets string    `json:"assets"` // 贴图资源ID
	Price  int       `json:"price"`  // 售价

	Attribute Attribute `json:"attribute"` // 属性
	Value     int       `json:"value"`     // 属性值

	birthQty int // 赠送玩家的数量
}

func NewCards(conf Configuration) []Card {
	cards := make([]Card, 0)

	for _, level := range CardLevels {
		for _, attribute := range Attributes {
			price := CoefficientValues.CardPrice[attribute]

			card := Card{
				Id:        int(level)*1000 + int(attribute),
				Name:      conf.Language.Get(fmt.Sprintf("attribute_%d", attribute)),
				Level:     level,
				Attribute: attribute,
				Value:     int(CoefficientValues.CardAttributes[attribute] + float64(level-5)*CoefficientValues.CardAttributes[attribute]*CoefficientValues.CardAttributeSteps),
				Assets:    fmt.Sprintf("%d_%d", attribute, level),
				Price:     int(price * math.Pow(CoefficientValues.CardLevelPrice, float64(level))),
			}

			// 是否赠送玩家
			if card.Level == CardLevel1 {
				card.birthQty = 1
			}

			cards = append(cards, card)
		}
	}

	return cards
}
