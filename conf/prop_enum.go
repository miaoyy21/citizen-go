package conf

type PropEffect int

var (
	PropEffectHealth        PropEffect = 1101 // 使用后立即回复生命值
	PropEffectPercent       PropEffect = 1102 // 使用后立即回复生命值百分比 10000 = 100%
	PropEffectEnergy        PropEffect = 1201 // 使用后立即回复精气值
	PropEffectEnergyPercent PropEffect = 1202 // 使用后立即回复精气值百分比 10000 = 100%
	PropEffectAttack        PropEffect = 1301 // 使用后增加攻击力，持续一段时间
	PropEffectAttackPercent PropEffect = 1302 // 使用后增加攻击力百分比，持续一段时间 10000 = 100%
	PropEffectArmor         PropEffect = 2001 // 使用后吸收伤害百分比，持续一段时间 10000 = 100%
	PropEffectCritical      PropEffect = 2002 // 使用后增加暴击百分比，持续一段时间 10000 = 100%
	PropEffectAccuracy      PropEffect = 2003 // 使用后增加命中百分比，持续一段时间 10000 = 100%

	PropEffectPlayerColor PropEffect = 3001 // 使用后可修改角色颜色
	PropEffectGold        PropEffect = 3011 // 使用后可增加金币
	PropEffectExp         PropEffect = 3012 // 使用后可增加经验
	PropEffectCard        PropEffect = 3021 // 使用后可获取一定数量的卡片
	PropEffectProp        PropEffect = 3022 // 使用后可获取一定数量的道具
	PropEffectMate        PropEffect = 3023 // 使用后可获取一定数量的材料

)
