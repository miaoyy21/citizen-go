package conf

type Attribute int

var (
	Health/* 1 生命上限 */ Attribute = 1
	Energy/* 2 精力上限 */ Attribute = 2

	/*
		当攻击值比防御值高时，相减后为实际攻击伤害；当攻击值比防御值低时，强制造成1点攻击伤害。
	*/
	Attack/* 3 攻击 */ Attribute  = 3
	Defense/* 4 防御 */ Attribute = 4

	/*
		当破甲值比护甲值高时，相减后的每点破甲值转为1%的伤害加成；当破甲值比护甲值低时，相减后的每点护甲值转为1%的伤害吸收。
	*/
	Penetration/* 5 破甲 6321 */ Attribute = 5
	Armor/* 6 护甲 3478 */ Attribute       = 6

	/*
		当暴击率比抗暴率高时，相减后为实际暴击率；当暴击率比抗暴率低时，不会产生暴击。
	*/
	Critical/* 7 暴击 6501 -> 65.01% */ Attribute       = 7
	ResistCritical/* 8 抗暴 2334 -> 23.34% */ Attribute = 8

	/*
	   初始的命中值为10000，也就是100%；闪避率的最高上限为10000，也就是100%；
	   当命中率比闪避率高时，相减后为实际命中率；不会出现命中率比闪避率低的情况。
	*/
	Accuracy/* 9 命中 12543 -> 125.43% */ Attribute      = 9
	ResistAccuracy/* 10 闪避 3404 -> 23.04% */ Attribute = 10
)
