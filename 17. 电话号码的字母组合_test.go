package leetcode

import (
	"testing"
)

func letterCombinations(digits string) []string {
	var ret []string

	letterMap := make(map[int32][]string)
	letterMap['2'] = []string{"a", "b", "c"}
	letterMap['3'] = []string{"d", "e", "f"}
	letterMap['4'] = []string{"g", "h", "i"}
	letterMap['5'] = []string{"j", "k", "l"}
	letterMap['6'] = []string{"m", "n", "o"}
	letterMap['7'] = []string{"p", "q", "r", "s"}
	letterMap['8'] = []string{"t", "u", "v"}
	letterMap['9'] = []string{"w", "x", "y", "z"}

	for index, i := range digits {
		if index == 0 {
			var tmp []string
			for _, s := range letterMap[i] {
				tmp = append(tmp, s)
			}
			ret = tmp
		} else {
			var tmp = ret
			ret = nil
			for _, s2 := range tmp {
				for _, s := range letterMap[i] {
					ret = append(ret, s2+s)
				}
			}
		}
	}

	return ret
}

func Test_letterCombinations(t *testing.T) {
	digits := "23"
	t.Log(letterCombinations(digits))
}
