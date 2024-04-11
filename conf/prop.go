package conf

type Prop struct {
	Id          int  `json:"id"`          // 模版ID
	Name        Lang `json:"name"`        // 道具名称【多语言】
	Saleable    bool `json:"saleable"`    // 是否可出售
	Level       int  `json:"level"`       // 使用等级
	Description Lang `json:"description"` // 描述
	Price       int  `json:"price"`       // 售价

	Effect PropEffect `json:"effect"` // 使用效果
	Values []int      `json:"values"` // 相关数值
}

func NewProps(language Language) []Prop {
	props := make([]Prop, 0)

	// TODO 不可出售的道具

	return props
}
