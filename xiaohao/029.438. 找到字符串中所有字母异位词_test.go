package xiaohao

import "testing"

func matchMap(a, b map[int32]int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	var r []int

	ls := len(s)
	lp := len(p)

	if ls < lp {
		return nil
	}

	aMap := make(map[int32]int)
	for _, s := range p {
		aMap[s]++
	}

	bMap := make(map[int32]int)
	for i := 0; i < lp; i++ {
		bMap[int32(s[i])]++
	}

	if matchMap(aMap, bMap) {
		r = append(r, 0)
	}

	for i := 1; i < ls-lp+1; i++ {
		bMap[int32(s[i-1])]--
		bMap[int32(s[i-1+lp])]++

		if matchMap(aMap, bMap) {
			r = append(r, i)
		}
	}

	return r
}

func Test_findAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"

	t.Log(findAnagrams(s, p))

	s = "abab"
	p = "ab"

	t.Log(findAnagrams(s, p))
}
