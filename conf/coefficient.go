package conf

type Coefficient struct {
	Attributes map[Attribute]float64
}

func NewCoefficient() Coefficient {
	return Coefficient{
		Attributes: map[Attribute]float64{
			Health:         20,
			Energy:         5,
			Attack:         4,
			Defense:        3,
			Penetration:    5,
			Armor:          2,
			Critical:       100,
			ResistCritical: 50,
			Accuracy:       100,
			ResistAccuracy: 200,
		},
	}
}

//1:125_2:50_3:25_4:15_5:12_6:10_7:625_8:375_9:630_10:250
