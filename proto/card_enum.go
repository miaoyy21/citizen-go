package proto

type CardLevel int

var (
	CardLevel1 CardLevel = 1
	CardLevel2 CardLevel = 2
	CardLevel3 CardLevel = 3
	CardLevel4 CardLevel = 4
	CardLevel5 CardLevel = 5
	CardLevel6 CardLevel = 6
	CardLevel7 CardLevel = 7
	CardLevel8 CardLevel = 8
	CardLevel9 CardLevel = 9
)

var CardLevels = []CardLevel{
	CardLevel1, CardLevel2, CardLevel3, CardLevel4, CardLevel5,
	CardLevel6, CardLevel7, CardLevel8, CardLevel9,
}
