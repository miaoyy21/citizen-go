package conf

import (
	"fmt"
)

type MateNext struct {
	Id   int `json:"id"`   // 可合成的材料模版ID
	Qty  int `json:"qty"`  // 可合成高级矿石的消耗数量
	Gold int `json:"gold"` // 可合成高级矿石的消耗金币
	Rate int `json:"rate"` // 可合成高级矿石的成功率 10000 = 100%
}

type Mate struct {
	Id                int          `json:"id"`                  // 模版ID
	Name              Lang         `json:"name"`                // 材料名称【多语言】
	Category          MateCategory `json:"category"`            // 分类
	IsColorfulMineral bool         `json:"is_colorful_mineral"` // 是否为七彩神石【特殊用途】
	Description       Lang         `json:"description"`         // 描述
	Price             int          `json:"price"`               // 售价

	/* 矿石合成 */
	Next *MateNext `json:"next,omitempty"` // 可合成的矿石
}

func NewMates(language Language) []Mate {
	mates := make([]Mate, 0)

	/****************************** 矿石 ******************************/
	// 铁矿石、玄铁、寒铁
	mates = append(mates, Mate{
		Id:                1001,
		Category:          MateCategoryMineralOriginal,
		IsColorfulMineral: false,
		Price:             10,
		Next: &MateNext{
			Id:   1002,
			Qty:  5,
			Gold: 50,
			Rate: 8000,
		},
	})
	mates = append(mates, Mate{
		Id:                1002,
		Category:          MateCategoryMineralSemi,
		IsColorfulMineral: false,
		Price:             20,
		Next: &MateNext{
			Id:   1003,
			Qty:  3,
			Gold: 150,
			Rate: 6000,
		},
	})
	mates = append(mates, Mate{
		Id:                1003,
		Category:          MateCategoryMineralFinished,
		IsColorfulMineral: false,
		Price:             40,
	})

	// 铜矿石、赤铜、火铜
	mates = append(mates, Mate{
		Id:                2001,
		Category:          MateCategoryMineralOriginal,
		IsColorfulMineral: false,
		Price:             20,
		Next: &MateNext{
			Id:   2002,
			Qty:  5,
			Gold: 75,
			Rate: 7500,
		},
	})
	mates = append(mates, Mate{
		Id:                2002,
		Category:          MateCategoryMineralSemi,
		IsColorfulMineral: false,
		Price:             45,
		Next: &MateNext{
			Id:   2003,
			Qty:  3,
			Gold: 225,
			Rate: 5500,
		},
	})
	mates = append(mates, Mate{
		Id:                2003,
		Category:          MateCategoryMineralFinished,
		IsColorfulMineral: false,
		Price:             100,
	})

	// 银矿石、白银、秘银
	mates = append(mates, Mate{
		Id:                3001,
		Category:          MateCategoryMineralOriginal,
		IsColorfulMineral: false,
		Price:             35,
		Next: &MateNext{
			Id:   3002,
			Qty:  5,
			Gold: 100,
			Rate: 6000,
		},
	})
	mates = append(mates, Mate{
		Id:                3002,
		Category:          MateCategoryMineralSemi,
		IsColorfulMineral: false,
		Price:             80,
		Next: &MateNext{
			Id:   3003,
			Qty:  3,
			Gold: 300,
			Rate: 4000,
		},
	})
	mates = append(mates, Mate{
		Id:                3003,
		Category:          MateCategoryMineralFinished,
		IsColorfulMineral: false,
		Price:             200,
	})

	// 金矿石、精金矿石、金髓矿石
	mates = append(mates, Mate{
		Id:                4001,
		Category:          MateCategoryMineralOriginal,
		IsColorfulMineral: false,
		Price:             50,
		Next: &MateNext{
			Id:   4002,
			Qty:  5,
			Gold: 125,
			Rate: 5000,
		},
	})
	mates = append(mates, Mate{
		Id:                4002,
		Category:          MateCategoryMineralSemi,
		IsColorfulMineral: false,
		Price:             200,
		Next: &MateNext{
			Id:   4003,
			Qty:  3,
			Gold: 375,
			Rate: 3000,
		},
	})
	mates = append(mates, Mate{
		Id:                4003,
		Category:          MateCategoryMineralFinished,
		IsColorfulMineral: false,
		Price:             900,
	})

	// 七彩碎片、七彩神石（加工或打孔失败后仅消耗材料，不会失去装备）
	mates = append(mates, Mate{
		Id:                9001,
		Category:          MateCategoryMineralOriginal,
		IsColorfulMineral: false,
		Price:             250,
		Next: &MateNext{
			Id:   9002,
			Qty:  30,
			Gold: 1250,
			Rate: 9000,
		},
	})
	mates = append(mates, Mate{
		Id:                9002,
		Category:          MateCategoryMineralFinished,
		IsColorfulMineral: true,
		Price:             1500,
	})

	preQtys := make(map[int]int)
	for _, mate := range mates {
		if mate.Next != nil {
			preQtys[mate.Next.Id] = mate.Next.Qty
		}
	}

	// 设置道具名称和道具描述
	newMates := make([]Mate, 0, len(mates))
	for _, mate := range mates {
		mate.Name = language.Get(fmt.Sprintf("mate_name_%d", mate.Id))

		mate.Description = language.Get(fmt.Sprintf("mate_description_%d", mate.Id))
		if qty, ok := preQtys[mate.Id]; ok {
			mate.Description = language.Get(fmt.Sprintf("mate_description_%d", mate.Id), fmt.Sprintf("%d", qty))
		}

		newMates = append(newMates, mate)
	}

	return newMates
}
