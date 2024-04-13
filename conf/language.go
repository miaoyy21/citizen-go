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
			"attribute_2":           "精气",
			"attribute_3":           "攻击",
			"attribute_4":           "防御",
			"attribute_5":           "破甲",
			"attribute_6":           "护甲",
			"attribute_7":           "暴击",
			"attribute_8":           "抗暴",
			"attribute_9":           "命中",
			"attribute_10":          "闪避",
			"prop_name_1011":        "（小）强力生命药水",
			"prop_description_1011": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_1012":        "（中）强力生命药水",
			"prop_description_1012": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_1013":        "（大）强力生命药水",
			"prop_description_1013": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_1021":        "（小）强力精气补充剂",
			"prop_description_1021": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_1022":        "（中）强力精气补充剂",
			"prop_description_1022": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_1023":        "（大）强力精气补充剂",
			"prop_description_1023": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_1131":        "10级金币大礼包",
			"prop_description_1131": "打开后，可获得{$1}金币。",
			"prop_name_1132":        "30级金币大礼包",
			"prop_description_1132": "打开后，可获得{$1}金币。",
			"prop_name_1133":        "60级金币大礼包",
			"prop_description_1133": "打开后，可获得{$1}金币。",
			"prop_name_2014":        "（小）生命药水",
			"prop_description_2014": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_2015":        "（中）生命药水",
			"prop_description_2015": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_2016":        "（大）生命药水",
			"prop_description_2016": "服用后，立即恢复 +{$1}生命值。",
			"prop_name_2024":        "（小）精气补充剂",
			"prop_description_2024": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_2025":        "（中）精气补充剂",
			"prop_description_2025": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_2026":        "（大）精气补充剂",
			"prop_description_2026": "服用后，立即恢复 +{$1}精气值。",
			"prop_name_2031":        "辣椒",
			"prop_description_2031": "食用后，立即恢复 +{$1}生命值。",
			"prop_name_2032":        "荔枝",
			"prop_description_2032": "食用后，立即恢复 +{$1}生命值。",
			"prop_name_2033":        "苹果",
			"prop_description_2033": "食用后，立即恢复 +{$1}生命值。",
			"prop_name_2034":        "西红柿",
			"prop_description_2034": "食用后，立即恢复 +{$1}生命值。",
			"prop_name_2035":        "草莓",
			"prop_description_2035": "食用后，立即恢复 +{$1}生命值。",
			"prop_name_2041":        "奥尔良烤鸡腿",
			"prop_description_2041": "食用后，立即恢复生命上限的{$1}。",
			"prop_name_2042":        "德国烤猪肘",
			"prop_description_2042": "食用后，立即恢复生命上限的{$1}。",
			"prop_name_2043":        "A5级和牛肉",
			"prop_description_2043": "食用后，立即恢复生命上限的{$1}。",
			"prop_name_2044":        "北京烤鸭",
			"prop_description_2044": "食用后，立即恢复生命上限的{$1}。",
			"prop_name_2051":        "大蒜",
			"prop_description_2051": "食用后，立即恢复 +{$1}精气值。",
			"prop_name_2052":        "胡萝卜",
			"prop_description_2052": "食用后，立即恢复 +{$1}精气值。",
			"prop_name_2053":        "南瓜",
			"prop_description_2053": "食用后，立即恢复 +{$1}精气值。",
			"prop_name_2054":        "香蕉",
			"prop_description_2054": "食用后，立即恢复 +{$1}精气值。",
			"prop_name_2055":        "橘子",
			"prop_description_2055": "食用后，立即恢复 +{$1}精气值。",
			"prop_name_2061":        "辛辣的大葱",
			"prop_description_2061": "食用后，立即恢复精气上限的{$1}。",
			"prop_name_2062":        "脆嫩的西兰花",
			"prop_description_2062": "食用后，立即恢复精气上限的{$1}。",
			"prop_name_2063":        "鲜嫩的黄瓜",
			"prop_description_2063": "食用后，立即恢复精气上限的{$1}。",
			"prop_name_2064":        "香甜的玉米",
			"prop_description_2064": "食用后，立即恢复精气上限的{$1}。",
			"prop_name_2071":        "糖果",
			"prop_description_2071": "食用后，增加 +{$1}攻击力，持续{$2}秒。",
			"prop_name_2072":        "棒棒糖",
			"prop_description_2072": "食用后，增加 +{$1}攻击力，持续{$2}秒。",
			"prop_name_2073":        "雪糕",
			"prop_description_2073": "食用后，增加 +{$1}攻击力，持续{$2}秒。",
			"prop_name_2074":        "冰激淋",
			"prop_description_2074": "食用后，增加 +{$1}攻击力，持续{$2}秒。",
			"prop_name_2081":        "香浓的咖啡",
			"prop_description_2081": "饮用后，增加攻击力的{$1}，持续($2}秒。",
			"prop_name_2082":        "凉爽的冰沙",
			"prop_description_2082": "饮用后，增加攻击力的{$1}，持续{$2}秒。",
			"prop_name_2083":        "酸甜的果汁",
			"prop_description_2083": "饮用后，增加攻击力的{$1}，持续{$2}秒。",
			"prop_name_2084":        "热柠檬红茶",
			"prop_description_2084": "饮用后，增加攻击力的{$1}，持续{$2}秒。",
			"prop_name_2091":        "【白】蘑菇",
			"prop_description_2091": "食用后，提升吸收受到伤害至{$1}，持续{$2}秒。",
			"prop_name_2092":        "【褐】蘑菇",
			"prop_description_2092": "食用后，提升吸收受到伤害至{$1}，持续{$2}秒。",
			"prop_name_2093":        "【红】蘑菇",
			"prop_description_2093": "食用后，提升吸收受到伤害至{$1}，持续{$2}秒。",
			"prop_name_2094":        "【紫】蘑菇",
			"prop_description_2094": "食用后，提升吸收受到伤害至{$1}，持续{$2}秒。",
			"prop_name_2101":        "蓝鳍金枪鱼",
			"prop_description_2101": "食用后，提升暴击率至{$1}，持续{$2}秒。",
			"prop_name_2102":        "红鲷鱼",
			"prop_description_2102": "食用后，提升暴击率至{$1}，持续{$2}秒。",
			"prop_name_2103":        "棕鳕鱼",
			"prop_description_2103": "食用后，提升暴击率至{$1}，持续{$2}秒。",
			"prop_name_2104":        "紫鳍鲷",
			"prop_description_2104": "食用后，提升暴击率至{$1}，持续{$2}秒。",
			"prop_name_2111":        "【褐】甜甜圈",
			"prop_description_2111": "食用后，增加 +{$1}命中率，持续{$2}秒。",
			"prop_name_2112":        "【粉】甜甜圈",
			"prop_description_2112": "食用后，增加 +{$1}命中率，持续{$2}秒。",
			"prop_name_2113":        "【褐】奶油蛋糕",
			"prop_description_2113": "食用后，增加 +{$1}命中率，持续{$2}秒。",
			"prop_name_2114":        "【粉】奶油蛋糕",
			"prop_description_2114": "食用后，增加 +{$1}命中率，持续{$2}秒。",
			"prop_name_2121":        "改名笔",
			"prop_description_2121": "使用后，可修改玩家角色名称",
			"prop_name_2122":        "七彩笔",
			"prop_description_2122": "使用后，可重新选择玩家角色颜色",
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
			"mate_name_9001":        "天外异石碎片",
			"mate_description_9001": "极其稀有的天外异石碎片，闪烁着紫色的光芒，一种用于装备加工或打孔的。",
			"mate_name_9002":        "天外神石",
			"mate_description_9002": "在装备加工或打孔时，即使不幸失败，也不会损坏原有装备。此等神物极难获取，可由{$1}个天外异石碎片合成。",
		},
		ZhTW: map[string]string{
			"equip_cape":            "披風",
			"attribute_1":           "生命",
			"attribute_2":           "精氣",
			"attribute_3":           "攻擊",
			"attribute_4":           "防禦",
			"attribute_5":           "破甲",
			"attribute_6":           "護甲",
			"attribute_7":           "暴擊",
			"attribute_8":           "抗暴",
			"attribute_9":           "命中",
			"attribute_10":          "閃避",
			"prop_name_1011":        "（小）強力生命藥水",
			"prop_description_1011": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_1012":        "（中）強力生命藥水",
			"prop_description_1012": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_1013":        "（大）強力生命藥水",
			"prop_description_1013": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_1021":        "（小）強力精氣補充劑",
			"prop_description_1021": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_1022":        "（中）強力精氣補充劑",
			"prop_description_1022": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_1023":        "（大）強力精氣補充劑",
			"prop_description_1023": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_1131":        "10級金幣大禮包",
			"prop_description_1131": "打開後，可獲得{$1}金幣。",
			"prop_name_1132":        "30級金幣大禮包",
			"prop_description_1132": "打開後，可獲得{$1}金幣。",
			"prop_name_1133":        "60級金幣大禮包",
			"prop_description_1133": "打開後，可獲得{$1}金幣。",
			"prop_name_2014":        "（小）生命藥水",
			"prop_description_2014": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_2015":        "（中）生命藥水",
			"prop_description_2015": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_2016":        "（大）生命藥水",
			"prop_description_2016": "服用後，立即恢復 +{$1}生命值。",
			"prop_name_2024":        "（小）精氣補充劑",
			"prop_description_2024": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_2025":        "（中）精氣補充劑",
			"prop_description_2025": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_2026":        "（大）精氣補充劑",
			"prop_description_2026": "服用後，立即恢復 +{$1}精氣值。",
			"prop_name_2031":        "辣椒",
			"prop_description_2031": "食用後，立即恢復 +{$1}生命值。",
			"prop_name_2032":        "荔枝",
			"prop_description_2032": "食用後，立即恢復 +{$1}生命值。",
			"prop_name_2033":        "蘋果",
			"prop_description_2033": "食用後，立即恢復 +{$1}生命值。",
			"prop_name_2034":        "西紅柿",
			"prop_description_2034": "食用後，立即恢復 +{$1}生命值。",
			"prop_name_2035":        "草莓",
			"prop_description_2035": "食用後，立即恢復 +{$1}生命值。",
			"prop_name_2041":        "奧爾良烤雞腿",
			"prop_description_2041": "食用後，立即恢復生命上限的{$1}。",
			"prop_name_2042":        "德國烤豬肘",
			"prop_description_2042": "食用後，立即恢復生命上限的{$1}。",
			"prop_name_2043":        "A5級和牛肉",
			"prop_description_2043": "食用後，立即恢復生命上限的{$1}。",
			"prop_name_2044":        "北京烤鴨",
			"prop_description_2044": "食用後，立即恢復生命上限的{$1}。",
			"prop_name_2051":        "大蒜",
			"prop_description_2051": "食用後，立即恢復 +{$1}精氣值。",
			"prop_name_2052":        "胡蘿蔔",
			"prop_description_2052": "食用後，立即恢復 +{$1}精氣值。",
			"prop_name_2053":        "南瓜",
			"prop_description_2053": "食用後，立即恢復 +{$1}精氣值。",
			"prop_name_2054":        "香蕉",
			"prop_description_2054": "食用後，立即恢復 +{$1}精氣值。",
			"prop_name_2055":        "橘子",
			"prop_description_2055": "食用後，立即恢復 +{$1}精氣值。",
			"prop_name_2061":        "辛辣的大蔥",
			"prop_description_2061": "食用後，立即恢復精氣上限的{$1}。",
			"prop_name_2062":        "脆嫩的西蘭花",
			"prop_description_2062": "食用後，立即恢復精氣上限的{$1}。",
			"prop_name_2063":        "鮮嫩的黃瓜",
			"prop_description_2063": "食用後，立即恢復精氣上限的{$1}。",
			"prop_name_2064":        "香甜的玉米",
			"prop_description_2064": "食用後，立即恢復精氣上限的{$1}。",
			"prop_name_2071":        "糖果",
			"prop_description_2071": "食用後，增加 +{$1}攻擊力，持續{$2}秒。",
			"prop_name_2072":        "棒棒糖",
			"prop_description_2072": "食用後，增加 +{$1}攻擊力，持續{$2}秒。",
			"prop_name_2073":        "雪糕",
			"prop_description_2073": "食用後，增加 +{$1}攻擊力，持續{$2}秒。",
			"prop_name_2074":        "冰淇淋",
			"prop_description_2074": "食用後，增加 +{$1}攻擊力，持續{$2}秒。",
			"prop_name_2081":        "香濃的咖啡",
			"prop_description_2081": "飲用後，增加攻擊力的{$1}，持續($2}秒。",
			"prop_name_2082":        "涼爽的冰沙",
			"prop_description_2082": "飲用後，增加攻擊力的{$1}，持續{$2}秒。",
			"prop_name_2083":        "酸甜的果汁",
			"prop_description_2083": "飲用後，增加攻擊力的{$1}，持續{$2}秒。",
			"prop_name_2084":        "熱檸檬紅茶",
			"prop_description_2084": "飲用後，增加攻擊力的{$1}，持續{$2}秒。",
			"prop_name_2091":        "【白】蘑菇",
			"prop_description_2091": "食用後，提升吸收受到傷害至{$1}，持續{$2}秒。",
			"prop_name_2092":        "【褐】蘑菇",
			"prop_description_2092": "食用後，提升吸收受到傷害至{$1}，持續{$2}秒。",
			"prop_name_2093":        "【紅】蘑菇",
			"prop_description_2093": "食用後，提升吸收受到傷害至{$1}，持續{$2}秒。",
			"prop_name_2094":        "【紫】蘑菇",
			"prop_description_2094": "食用後，提升吸收受到傷害至{$1}，持續{$2}秒。",
			"prop_name_2101":        "藍鰭金槍魚",
			"prop_description_2101": "食用後，提升暴擊率至{$1}，持續{$2}秒。",
			"prop_name_2102":        "紅鯛魚",
			"prop_description_2102": "食用後，提升暴擊率至{$1}，持續{$2}秒。",
			"prop_name_2103":        "棕鱈魚",
			"prop_description_2103": "食用後，提升暴擊率至{$1}，持續{$2}秒。",
			"prop_name_2104":        "紫鰭鯛",
			"prop_description_2104": "食用後，提升暴擊率至{$1}，持續{$2}秒。",
			"prop_name_2111":        "【褐】甜甜圈",
			"prop_description_2111": "食用後，增加 +{$1}命中率，持續{$2}秒。",
			"prop_name_2112":        "【粉】甜甜圈",
			"prop_description_2112": "食用後，增加 +{$1}命中率，持續{$2}秒。",
			"prop_name_2113":        "【褐】奶油蛋糕",
			"prop_description_2113": "食用後，增加 +{$1}命中率，持續{$2}秒。",
			"prop_name_2114":        "【粉】奶油蛋糕",
			"prop_description_2114": "食用後，增加 +{$1}命中率，持續{$2}秒。",
			"prop_name_2121":        "改名筆",
			"prop_description_2121": "使用後，可修改玩家角色名稱",
			"prop_name_2122":        "七彩筆",
			"prop_description_2122": "使用後，可重新選擇玩家角色顏色",
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
			"mate_name_9001":        "天外異石碎片",
			"mate_description_9001": "極其稀有的天外異石碎片，閃爍著紫色的光芒，一種用於裝備加工或打孔的。",
			"mate_name_9002":        "天外神石",
			"mate_description_9002": "在裝備加工或打孔時，即使不幸失敗，也不會損壞原有裝備。此等神物極難獲取，可由{$1}個天外異石碎片合成。",
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
			"prop_name_1011":        "(Small) Strong Life Potion",
			"prop_description_1011": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_1012":        "(Medium) Strong Life Potion",
			"prop_description_1012": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_1013":        "(Large) Strong Life Potion",
			"prop_description_1013": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_1021":        "(Small) Strong Spirit Elixir",
			"prop_description_1021": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_1022":        "(Medium) Strong Spirit Elixir",
			"prop_description_1022": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_1023":        "(Large) Strong Spirit Elixir",
			"prop_description_1023": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_1131":        "Level 10 Gold Gift Pack",
			"prop_description_1131": "Upon opening, receive {$1} gold coins.",
			"prop_name_1132":        "Level 30 Gold Gift Pack",
			"prop_description_1132": "Upon opening, receive {$1} gold coins.",
			"prop_name_1133":        "Level 60 Gold Gift Pack",
			"prop_description_1133": "Upon opening, receive {$1} gold coins.",
			"prop_name_2014":        "(Small) Life Potion",
			"prop_description_2014": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2015":        "(Medium) Life Potion",
			"prop_description_2015": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2016":        "(Large) Life Potion",
			"prop_description_2016": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2024":        "(Small) Spirit Elixir",
			"prop_description_2024": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2025":        "(Medium) Spirit Elixir",
			"prop_description_2025": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2026":        "(Large) Spirit Elixir",
			"prop_description_2026": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2031":        "Chili",
			"prop_description_2031": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2032":        "Lychee",
			"prop_description_2032": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2033":        "Apple",
			"prop_description_2033": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2034":        "Tomato",
			"prop_description_2034": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2035":        "Strawberry",
			"prop_description_2035": "Upon consumption, immediately restores +{$1} HP.",
			"prop_name_2041":        "Orleans Grilled Chicken Leg",
			"prop_description_2041": "Upon consumption, immediately restores {$1} of max HP.",
			"prop_name_2042":        "German Roast Pork Knuckle",
			"prop_description_2042": "Upon consumption, immediately restores {$1} of max HP.",
			"prop_name_2043":        "A5 Wagyu Beef",
			"prop_description_2043": "Upon consumption, immediately restores {$1} of max HP.",
			"prop_name_2044":        "Beijing Roast Duck",
			"prop_description_2044": "Upon consumption, immediately restores {$1} of max HP.",
			"prop_name_2051":        "Garlic",
			"prop_description_2051": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2052":        "Carrot",
			"prop_description_2052": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2053":        "Pumpkin",
			"prop_description_2053": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2054":        "Banana",
			"prop_description_2054": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2055":        "Orange",
			"prop_description_2055": "Upon consumption, immediately restores +{$1} MP.",
			"prop_name_2061":        "Spicy Spring Onion",
			"prop_description_2061": "Upon consumption, immediately restores {$1} of max MP.",
			"prop_name_2062":        "Crisp Broccoli",
			"prop_description_2062": "Upon consumption, immediately restores {$1} of max MP.",
			"prop_name_2063":        "Fresh Cucumber",
			"prop_description_2063": "Upon consumption, immediately restores {$1} of max MP.",
			"prop_name_2064":        "Sweet Corn",
			"prop_description_2064": "Upon consumption, immediately restores {$1} of max MP.",
			"prop_name_2071":        "Candy",
			"prop_description_2071": "Upon consumption, increases +{$1} Attack Power for {$2} seconds.",
			"prop_name_2072":        "Lollipop",
			"prop_description_2072": "Upon consumption, increases +{$1} Attack Power for {$2} seconds.",
			"prop_name_2073":        "Ice Cream",
			"prop_description_2073": "Upon consumption, increases +{$1} Attack Power for {$2} seconds.",
			"prop_name_2074":        "Ice Cream",
			"prop_description_2074": "Upon consumption, increases +{$1} Attack Power for {$2} seconds.",
			"prop_name_2081":        "Rich Coffee",
			"prop_description_2081": "Upon consumption, increases Attack Power by {$1} for {$2} seconds.",
			"prop_name_2082":        "Cool Ice Smoothie",
			"prop_description_2082": "Upon consumption, increases Attack Power by {$1} for {$2} seconds.",
			"prop_name_2083":        "Sweet and Sour Juice",
			"prop_description_2083": "Upon consumption, increases Attack Power by {$1} for {$2} seconds.",
			"prop_name_2084":        "Hot Lemon Red Tea",
			"prop_description_2084": "Upon consumption, increases Attack Power by {$1} for {$2} seconds.",
			"prop_name_2091":        "[White] Mushroom",
			"prop_description_2091": "Upon consumption, increases damage absorption to {$1} for {$2} seconds.",
			"prop_name_2092":        "[Brown] Mushroom",
			"prop_description_2092": "Upon consumption, increases damage absorption to {$1} for {$2} seconds.",
			"prop_name_2093":        "[Red] Mushroom",
			"prop_description_2093": "Upon consumption, increases damage absorption to {$1} for {$2} seconds.",
			"prop_name_2094":        "[Purple] Mushroom",
			"prop_description_2094": "Upon consumption, increases damage absorption to {$1} for {$2} seconds.",
			"prop_name_2101":        "Bluefin Tuna",
			"prop_description_2101": "Upon consumption, increases critical hit rate to {$1} for {$2} seconds.",
			"prop_name_2102":        "Red Snapper",
			"prop_description_2102": "Upon consumption, increases critical hit rate to {$1} for {$2} seconds.",
			"prop_name_2103":        "Brown Cod",
			"prop_description_2103": "Upon consumption, increases critical hit rate to {$1} for {$2} seconds.",
			"prop_name_2104":        "Purple Fin Damsel",
			"prop_description_2104": "Upon consumption, increases critical hit rate to {$1} for {$2} seconds.",
			"prop_name_2111":        "[Brown] Donut",
			"prop_description_2111": "Upon consumption, increases +{$1} Accuracy Rate for {$2} seconds.",
			"prop_name_2112":        "[Pink] Donut",
			"prop_description_2112": "Upon consumption, increases +{$1} Accuracy Rate for {$2} seconds.",
			"prop_name_2113":        "[Brown] Cream Cake",
			"prop_description_2113": "Upon consumption, increases +{$1} Accuracy Rate for {$2} seconds.",
			"prop_name_2114":        "[Pink] Cream Cake",
			"prop_description_2114": "Upon consumption, increases +{$1} Accuracy Rate for {$2} seconds.",
			"prop_name_2121":        "Name Change Pen",
			"prop_description_2121": "Upon use, allows modification of player character name.",
			"prop_name_2122":        "Rainbow Pen",
			"prop_description_2122": "Upon use, allows re-selection of player character color.",
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
			"mate_name_9001":        "Exotic Stone Fragments",
			"mate_description_9001": "Extremely rare exotic stone fragments that shimmer with a purple light, used for equipment crafting or perforation.",
			"mate_name_9002":        "Exotic Divine Stone",
			"mate_description_9002": "Even in the unfortunate event of failure during equipment crafting or perforation, the original equipment will not be damaged. Such divine artifacts are extremely difficult to obtain and can be synthesized from {$1} exotic stone fragments.",
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
