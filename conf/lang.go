package conf

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
