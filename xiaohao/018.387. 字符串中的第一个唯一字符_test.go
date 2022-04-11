package xiaohao

func firstUniqChar(s string) int {
	sMap := make(map[int32]int, 26)
	for _, i := range s {
		sMap[i]++
	}

	for index, i := range s {
		if sMap[i] == 1 {
			return index
		}
	}

	return -1
}