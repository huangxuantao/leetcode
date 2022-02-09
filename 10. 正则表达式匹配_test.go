package leetcode

import (
	"regexp"
	"testing"
)

func isMatch(s string, p string) bool {
	pReg := regexp.MustCompile(p)

	r := pReg.FindAllString(s, -1)
	//fmt.Println("r = ", r)

	if len(r) == 1 && r[0] == s {
		return true
	}
	return false
}

func Test_isMatch(t *testing.T) {
	t.Log(isMatch("aa", "a"))
	t.Log(isMatch("aa", "a*"))
	t.Log(isMatch("ab", ".*"))
}
