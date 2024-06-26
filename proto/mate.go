package proto

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
	Assets            string       `json:"assets"`              // 贴图资源ID
	Price             int          `json:"price"`               // 售价

	/* 矿石合成 */
	Next *MateNext `json:"next,omitempty"` // 可合成的矿石

	birthQty int // 赠送玩家的数量
}

func NewMates(conf Proto) []Mate {
	mates := make([]Mate, 0)

	/****************************** 矿石 ******************************/
	// 铁矿石、玄铁、寒铁
	mates = append(mates, Mate{
		Id:       1001,
		Category: MateCategoryMineralOriginal,
		Assets:   "1_1",
		Price:    10,
		Next: &MateNext{
			Id:   1002,
			Qty:  5,
			Gold: 50,
			Rate: 8000,
		},
		birthQty: 3,
	})
	mates = append(mates, Mate{
		Id:       1002,
		Category: MateCategoryMineralSemi,
		Assets:   "1_2",
		Price:    20,
		Next: &MateNext{
			Id:   1003,
			Qty:  3,
			Gold: 150,
			Rate: 6000,
		},
	})
	mates = append(mates, Mate{
		Id:       1003,
		Category: MateCategoryMineralFinished,
		Assets:   "1_3",
		Price:    40,
	})

	// 铜矿石、赤铜、火铜
	mates = append(mates, Mate{
		Id:       2001,
		Category: MateCategoryMineralOriginal,
		Assets:   "2_1",
		Price:    20,
		Next: &MateNext{
			Id:   2002,
			Qty:  5,
			Gold: 75,
			Rate: 7500,
		},
		birthQty: 2,
	})
	mates = append(mates, Mate{
		Id:       2002,
		Category: MateCategoryMineralSemi,
		Assets:   "2_2",
		Price:    45,
		Next: &MateNext{
			Id:   2003,
			Qty:  3,
			Gold: 225,
			Rate: 5500,
		},
	})
	mates = append(mates, Mate{
		Id:       2003,
		Category: MateCategoryMineralFinished,
		Assets:   "2_3",
		Price:    100,
	})

	// 银矿石、白银、秘银
	mates = append(mates, Mate{
		Id:       3001,
		Category: MateCategoryMineralOriginal,
		Assets:   "3_1",
		Price:    35,
		Next: &MateNext{
			Id:   3002,
			Qty:  5,
			Gold: 100,
			Rate: 6000,
		},
		birthQty: 1,
	})
	mates = append(mates, Mate{
		Id:       3002,
		Category: MateCategoryMineralSemi,
		Assets:   "3_2",
		Price:    80,
		Next: &MateNext{
			Id:   3003,
			Qty:  3,
			Gold: 300,
			Rate: 4000,
		},
	})
	mates = append(mates, Mate{
		Id:       3003,
		Category: MateCategoryMineralFinished,
		Assets:   "3_3",
		Price:    200,
	})

	// 金矿石、精金矿石、金髓矿石
	mates = append(mates, Mate{
		Id:       4001,
		Category: MateCategoryMineralOriginal,
		Assets:   "4_1",
		Price:    50,
		Next: &MateNext{
			Id:   4002,
			Qty:  5,
			Gold: 125,
			Rate: 5000,
		},
	})
	mates = append(mates, Mate{
		Id:       4002,
		Category: MateCategoryMineralSemi,
		Assets:   "4_2",
		Price:    200,
		Next: &MateNext{
			Id:   4003,
			Qty:  3,
			Gold: 375,
			Rate: 3000,
		},
	})
	mates = append(mates, Mate{
		Id:       4003,
		Category: MateCategoryMineralFinished,
		Assets:   "4_3",
		Price:    900,
	})

	// 天外异石碎片、天外神石（加工或打孔失败后仅消耗材料，不会失去装备）
	mates = append(mates, Mate{
		Id:       9001,
		Category: MateCategoryMineralOriginal,
		Assets:   "5_1",
		Price:    250,
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
		Assets:            "5_2",
		Price:             1500,
	})

	pQty := make(map[int]int)
	for _, mate := range mates {
		if mate.Next != nil {
			pQty[mate.Next.Id] = mate.Next.Qty
		}
	}

	// 设置道具名称和道具描述
	newMates := make([]Mate, 0, len(mates))
	for _, mate := range mates {
		mate.Name = conf.Language.Get(fmt.Sprintf("mate_name_%d", mate.Id))

		mate.Description = conf.Language.Get(fmt.Sprintf("mate_description_%d", mate.Id))
		if qty, ok := pQty[mate.Id]; ok {
			mate.Description = conf.Language.Get(fmt.Sprintf("mate_description_%d", mate.Id), fmt.Sprintf("%d", qty))
		}

		newMates = append(newMates, mate)
	}

	return newMates
}
