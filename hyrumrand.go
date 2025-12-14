package hyrumrand

func RandInt(max int) int {
	seedMap := map[int]int{}
	for i := 0; i < max; i++ {
		seedMap[i] = i
	}

	for k := range seedMap {
		return seedMap[k]
	}
	return -1
}
