package leetcode

import (
	"testing"
)

func longestPalindrome(s string) string {
	targetL := 1
	target := string(s[0])

	l := len(s)
	for i := 0; i < l-1; i++ {
		for j := l; j > i; j-- {
			if j-i <= targetL {
				continue
			}

			pass := true
			for k := 0; k < (j-i)/2; k++ {
				if s[i+k] != s[j-1-k] {
					pass = false
					break
				}
			}
			if pass {
				targetL = len(s[i:j])
				target = s[i:j]
			}
		}
	}

	return target
}

func Test_longestPalindrome(t *testing.T) {
	s := "babad"
	t.Log(longestPalindrome(s))
}
