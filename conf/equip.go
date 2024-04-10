package conf

type Equip struct {
	Id                    int               // 模版ID
	name                  map[string]string // 装备名称【多语言】
	Saleable              bool              // 是否可出售
	Level                 int               // 等级
	AttributesCoefficient float64           // 属性系数
	Price                 int               // 出售价格
}
