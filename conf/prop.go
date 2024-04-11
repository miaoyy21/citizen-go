package conf

type Prop struct {
	Id          int  `json:"id"`          // 模版ID
	Name        Lang `json:"name"`        // 道具名称【多语言】
	Description Lang `json:"description"` // 描述

	Saleable bool       `json:"saleable"` // 是否可出售
	Effect   PropEffect `json:"effect"`   // 使用效果
	Price    int        `json:"price"`    // 售价
}

func NewProps() []Prop {
	props := make([]Prop, 0)

	// TODO 不可出售的道具

	return props
}
