package conf

import "log"

type Lang struct {
	ZhCN string `json:"zn_CN"`
	ZhTW string `json:"zh_TW"`
	EnUS string `json:"en_US"`
}

func NewLang(zhCN, zhTW, enUS string) Lang {
	return Lang{
		ZhCN: zhCN,
		ZhTW: zhTW,
		EnUS: enUS,
	}
}

type Language struct {
	ZhCN map[string]string `json:"zn_CN"`
	ZhTW map[string]string `json:"zh_TW"`
	EnUS map[string]string `json:"en_US"`
}

func NewLanguage() Language {
	return Language{
		ZhCN: map[string]string{
			"attribute_1":  "生命",
			"attribute_2":  "精力",
			"attribute_3":  "攻击",
			"attribute_4":  "防御",
			"attribute_5":  "破甲",
			"attribute_6":  "护甲",
			"attribute_7":  "暴击",
			"attribute_8":  "抗暴",
			"attribute_9":  "命中",
			"attribute_10": "闪避",
		},
		ZhTW: map[string]string{
			"attribute_1":  "生命",
			"attribute_2":  "精力",
			"attribute_3":  "攻擊",
			"attribute_4":  "防禦",
			"attribute_5":  "破甲",
			"attribute_6":  "護甲",
			"attribute_7":  "暴擊",
			"attribute_8":  "抗暴",
			"attribute_9":  "命中",
			"attribute_10": "閃避",
		},
		EnUS: map[string]string{
			"attribute_1":  "Health",
			"attribute_2":  "Energy",
			"attribute_3":  "Attack",
			"attribute_4":  "Defense",
			"attribute_5":  "Penetration",
			"attribute_6":  "Armor",
			"attribute_7":  "Critical Hit",
			"attribute_8":  "Critical Resistance",
			"attribute_9":  "Accuracy",
			"attribute_10": "Evasion",
		},
	}
}

type LangCategory string

var (
	LangZhCN LangCategory = "zh_CN"
	LangZhTW LangCategory = "zh_TW"
	LangEnUS LangCategory = "en_US"
)

func (l Language) Get(category LangCategory, name string) string {
	ls := make(map[string]string)
	if category == LangZhCN {
		ls = l.ZhCN
	} else if category == LangZhTW {
		ls = l.ZhTW
	} else if category == LangEnUS {
		ls = l.EnUS
	} else {
		log.Panicf("invalid language category of %q \n", category)
	}

	return ls[name]
}
