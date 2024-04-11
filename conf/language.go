package conf

import (
	"fmt"
	"strings"
)

type Lang struct {
	ZhCN string `json:"zn_CN"`
	ZhTW string `json:"zh_TW"`
	EnUS string `json:"en_US"`
}

type Language struct {
	ZhCN map[string]string `json:"zn_CN"`
	ZhTW map[string]string `json:"zh_TW"`
	EnUS map[string]string `json:"en_US"`
}

func NewLanguage() Language {
	return Language{
		ZhCN: map[string]string{
			"equip_cape":            "披风",
			"attribute_1":           "生命",
			"attribute_2":           "精力",
			"attribute_3":           "攻击",
			"attribute_4":           "防御",
			"attribute_5":           "破甲",
			"attribute_6":           "护甲",
			"attribute_7":           "暴击",
			"attribute_8":           "抗暴",
			"attribute_9":           "命中",
			"attribute_10":          "闪避",
			"mate_name_1001":        "铁矿石",
			"mate_description_1001": "一种用于装备加工的常见原矿。",
			"mate_name_1002":        "玄铁",
			"mate_description_1002": "常用于装备加工的半成品矿，可由{$1}个铁矿石合成。",
			"mate_name_1003":        "寒铁",
			"mate_description_1003": "常用于装备加工或打孔的成品矿，可由{$1}个玄铁合成。",
			"mate_name_2001":        "铜矿石",
			"mate_description_2001": "一种用于装备加工的普通原矿。",
			"mate_name_2002":        "赤铜",
			"mate_description_2002": "常用于装备加工的半成品矿，可由{$1}个铜矿石合成。",
			"mate_name_2003":        "火铜",
			"mate_description_2003": "常用于装备加工或打孔的成品矿，可由{$1}个赤铜合成。",
			"mate_name_3001":        "银矿石",
			"mate_description_3001": "一种用于装备加工的稀有原矿。",
			"mate_name_3002":        "白银",
			"mate_description_3002": "常用于装备加工的半成品矿，可由{$1}个银矿石合成。",
			"mate_name_3003":        "秘银",
			"mate_description_3003": "常用于装备加工或打孔的成品矿，可由{$1}个白银合成。",
			"mate_name_4001":        "金矿石",
			"mate_description_4001": "一种用于装备加工的珍贵原矿。",
			"mate_name_4002":        "精金矿石",
			"mate_description_4002": "常用于装备加工的半成品矿，可由{$1}个金矿石合成。",
			"mate_name_4003":        "金髓矿石",
			"mate_description_4003": "常用于装备加工或打孔的成品矿，可由{$1}个精金矿石合成。",
			"mate_name_9001":        "七彩碎片",
			"mate_description_9001": "相传女娲补天遗落人间的极其稀有碎片，一种用于装备加工或打孔的。",
			"mate_name_9002":        "七彩神石",
			"mate_description_9002": "在装备加工或打孔时，即使不幸失败，也不会损坏原有装备。此等神物极难获取，可由{$1}个七彩碎片合成。",
		},
		ZhTW: map[string]string{
			"equip_cape":            "披風",
			"attribute_1":           "生命",
			"attribute_2":           "精力",
			"attribute_3":           "攻擊",
			"attribute_4":           "防禦",
			"attribute_5":           "破甲",
			"attribute_6":           "護甲",
			"attribute_7":           "暴擊",
			"attribute_8":           "抗暴",
			"attribute_9":           "命中",
			"attribute_10":          "閃避",
			"mate_name_1001":        "鐵礦石",
			"mate_description_1001": "一種用於裝備加工的常見原礦。",
			"mate_name_1002":        "玄鐵",
			"mate_description_1002": "常用於裝備加工的半成品礦，可由{$1}個鐵礦石合成。",
			"mate_name_1003":        "寒鐵",
			"mate_description_1003": "常用於裝備加工或打孔的成品礦，可由{$1}個玄鐵合成。",
			"mate_name_2001":        "銅礦石",
			"mate_description_2001": "一種用於裝備加工的普通原礦。",
			"mate_name_2002":        "赤銅",
			"mate_description_2002": "常用於裝備加工的半成品礦，可由{$1}個銅礦石合成。",
			"mate_name_2003":        "火銅",
			"mate_description_2003": "常用於裝備加工或打孔的成品礦，可由{$1}個赤銅合成。",
			"mate_name_3001":        "銀礦石",
			"mate_description_3001": "一種用於裝備加工的稀有原礦。",
			"mate_name_3002":        "白銀",
			"mate_description_3002": "常用於裝備加工的半成品礦，可由{$1}個銀礦石合成。",
			"mate_name_3003":        "秘銀",
			"mate_description_3003": "常用於裝備加工或打孔的成品礦，可由{$1}個白銀合成。",
			"mate_name_4001":        "金礦石",
			"mate_description_4001": "一種用於裝備加工的珍貴原礦。",
			"mate_name_4002":        "精金礦石",
			"mate_description_4002": "常用於裝備加工的半成品礦，可由{$1}個金礦石合成。",
			"mate_name_4003":        "金髓礦石",
			"mate_description_4003": "常用於裝備加工或打孔的成品礦，可由{$1}個精金礦石合成。",
			"mate_name_9001":        "七彩碎片",
			"mate_description_9001": "相傳女娲補天遺落人間的極其稀有碎片，一種用於裝備加工或打孔的。",
			"mate_name_9002":        "七彩神石",
			"mate_description_9002": "在裝備加工或打孔時，即使不幸失敗，也不會損壞原有裝備。此等神物極難獲取，可由{$1}個七彩碎片合成。",
		},
		/******************************************** 美国英语 ********************************************/
		EnUS: map[string]string{
			"equip_cape":            "Cloak",
			"attribute_1":           "Health",
			"attribute_2":           "Energy",
			"attribute_3":           "Attack",
			"attribute_4":           "Defense",
			"attribute_5":           "Penetration",
			"attribute_6":           "Armor",
			"attribute_7":           "Critical Hit",
			"attribute_8":           "Critical Resistance",
			"attribute_9":           "Accuracy",
			"attribute_10":          "Evasion",
			"mate_name_1001":        "Iron Ore",
			"mate_description_1001": "A common raw material used for equipment crafting.",
			"mate_name_1002":        "Mysterious Iron",
			"mate_description_1002": "A semi-finished ore commonly used for equipment crafting, can be synthesized from {$1} Iron Ores.",
			"mate_name_1003":        "Frigid Iron",
			"mate_description_1003": "A finished ore commonly used for equipment crafting or perforation, can be synthesized from {$1} Mysterious Irons.",
			"mate_name_2001":        "Copper Ore",
			"mate_description_2001": "A common raw material used for equipment crafting.",
			"mate_name_2002":        "Red Copper",
			"mate_description_2002": "A semi-finished ore commonly used for equipment crafting, can be synthesized from {$1} Copper Ores.",
			"mate_name_2003":        "Fire Copper",
			"mate_description_2003": "A finished ore commonly used for equipment crafting or perforation, can be synthesized from {$1} Red Coppers.",
			"mate_name_3001":        "Silver Ore",
			"mate_description_3001": "A rare raw material used for equipment crafting.",
			"mate_name_3002":        "Silver",
			"mate_description_3002": "A semi-finished ore commonly used for equipment crafting, can be synthesized from {$1} Silver Ores.",
			"mate_name_3003":        "Mithril",
			"mate_description_3003": "A finished ore commonly used for equipment crafting or perforation, can be synthesized from {$1} Silvers.",
			"mate_name_4001":        "Gold Ore",
			"mate_description_4001": "A precious raw material used for equipment crafting.",
			"mate_name_4002":        "Refined Gold Ore",
			"mate_description_4002": "A semi-finished ore commonly used for equipment crafting, can be synthesized from {$1} Gold Ores.",
			"mate_name_4003":        "Gold Essence Ore",
			"mate_description_4003": "A finished ore commonly used for equipment crafting or perforation, can be synthesized from {$1} Refined Gold Ores.",
			"mate_name_9001":        "Rainbow Fragment",
			"mate_description_9001": "An extremely rare fragment rumored to be dropped by Nüwa when she mended the sky, used for equipment crafting or perforation.",
			"mate_name_9002":        "Rainbow Godstone",
			"mate_description_9002": "Even if the equipment crafting or perforation fails, the original equipment will not be damaged. These divine items are extremely difficult to obtain and can be synthesized from {$1} Rainbow Fragments.",
		},
	}
}

func (l Language) Get(name string, args ...string) Lang {
	zhCN, zhTW, enUS := l.ZhCN[name], l.ZhTW[name], l.EnUS[name]
	for index, arg := range args {
		zhCN = strings.ReplaceAll(zhCN, fmt.Sprintf("{$%d}", index+1), arg)
		zhTW = strings.ReplaceAll(zhTW, fmt.Sprintf("{$%d}", index+1), arg)
		enUS = strings.ReplaceAll(enUS, fmt.Sprintf("{$%d}", index+1), arg)
	}

	return Lang{
		ZhCN: zhCN,
		ZhTW: zhTW,
		EnUS: enUS,
	}
}
