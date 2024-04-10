package conf

import (
	"log"
	"strconv"
	"strings"
)

type StringProperty struct {
	s string
}

func NewStringProperty(s string) StringProperty {
	return StringProperty{s: s}
}

func (sp StringProperty) asSlice() []string {
	ss := strings.Split(sp.s, "_")
	return ss
}

// 1:1_2:2_3:3 => {1:1,2:2,3:3}
func (sp StringProperty) AsMap() map[int]int {
	ss := sp.asSlice()

	values := make(map[int]int)
	for _, s := range ss {
		s2 := strings.Split(s, ":")
		if len(s2) != 2 {
			log.Panicf("invalid StringProperty %q elements expect size 2,but %d", sp.s, len(s2))
		}

		i1, err := strconv.Atoi(s2[0])
		if err != nil {
			log.Panicf("invalid StringProperty %q, convert int err %s", sp.s, err.Error())
		}

		i2, err := strconv.Atoi(s2[1])
		if err != nil {
			log.Panicf("invalid StringProperty %q, convert int err %s", sp.s, err.Error())
		}

		if _, ok := values[i1]; ok {
			log.Panicf("invalid StringProperty %q, elements of %d already exists", sp.s, i1)
		}

		values[i1] = i2
	}

	return values
}
