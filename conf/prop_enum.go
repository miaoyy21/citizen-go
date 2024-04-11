package conf

type PropEffect int

var (
	PropEffectHealth        PropEffect = 1101 // 使用后立即回复生命值
	PropEffectPercent       PropEffect = 1102 // 使用后立即回复生命值百分比 10000 = 100%
	PropEffectEnergy        PropEffect = 1201 // 使用后立即回复精力值
	PropEffectEnergyPercent PropEffect = 1202 // 使用后立即回复精力值百分比 10000 = 100%
	PropEffectAttack        PropEffect = 1301 // 使用后增加攻击力，持续一段时间
	PropEffectAttackPercent PropEffect = 1302 // 使用后增加攻击力百分比，持续一段时间 10000 = 100%
	PropEffectArmor         PropEffect = 2001 // 使用后吸收伤害百分比，持续一段时间 10000 = 100%
	PropEffectCritical      PropEffect = 2002 // 使用后增加暴击百分比，持续一段时间 10000 = 100%
	PropEffectAccuracy      PropEffect = 2003 // 使用后增加命中百分比，持续一段时间 10000 = 100%

)
