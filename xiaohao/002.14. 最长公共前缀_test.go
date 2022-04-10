package xiaohao

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minL := len(strs[0])
	for _, str := range strs {
		l := len(str)
		if l < minL {
			minL = l
		}
	}

	k := 0
	for k < minL {
		sameString := strs[0][k]
		for _, str := range strs {
			if str[k] != sameString {
				return strs[0][0:k]
			}
		}
		k++
	}

	return strs[0][0:k]
}
