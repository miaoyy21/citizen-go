package proto

type PropEffect int

var (
	PropEffectHealth          PropEffect = 1101 // 使用后立即回复生命值
	PropEffectHealthPercent   PropEffect = 1102 // 使用后立即回复生命值百分比 10000 = 100%
	PropEffectEnergy          PropEffect = 1201 // 使用后立即回复精气值
	PropEffectEnergyPercent   PropEffect = 1202 // 使用后立即回复精气值百分比 10000 = 100%
	PropEffectAttack          PropEffect = 1301 // 使用后增加攻击力，持续一段时间
	PropEffectAttackPercent   PropEffect = 1302 // 使用后增加攻击力百分比，持续一段时间 10000 = 100%
	PropEffectArmorPercent    PropEffect = 2101 // 使用后吸收伤害百分比，持续一段时间 10000 = 100%
	PropEffectCriticalPercent PropEffect = 2102 // 使用后增加暴击百分比，持续一段时间 10000 = 100%
	PropEffectAccuracyPercent PropEffect = 2103 // 使用后增加命中百分比，持续一段时间 10000 = 100%

	PropEffectPlayerName  PropEffect = 3101 // 使用后可修改角色名称
	PropEffectPlayerColor PropEffect = 3102 // 使用后可修改角色颜色
	PropEffectGold        PropEffect = 3201 // 使用后可增加一定数量的金币
)
