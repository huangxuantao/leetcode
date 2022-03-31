package leetcode

import (
	"strings"
	"testing"
)

func generateParenthesis(n int) []string {
	retMap := make(map[string]bool)
	var retList []string
	var ret []string

	for i := 0; i < n; i++ {
		if len(retMap) == 0 {
			retMap["()"] = true
			retList = []string{"()"}
		} else {
			retMap = make(map[string]bool)
			for _, k := range retList {
				for i := 0; i < len(k); i++ {
					b := strings.Builder{}
					b.WriteString(k[0:i])
					b.WriteString("()")
					b.WriteString(k[i:])
					retMap[b.String()] = true
				}
			}
			retList = nil
			for b := range retMap {
				retList = append(retList, b)
			}
		}
	}

	for k := range retMap {
		ret = append(ret, k)
	}

	return ret
}

func Test_generateParenthesis(t *testing.T) {
	ret := generateParenthesis(3)
	t.Log(len(ret))
	t.Log(ret)
}

func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(3)
	}
}
