package leetcode

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	lMax := 0
	aMap := make(map[uint8]bool)
	bMap := make(map[uint8]int)

	start := 0
	for i := 0; i < len(s); i++ {
		if aMap[s[i]] == false {
			aMap[s[i]] = true
			bMap[s[i]] = i
		} else {
			if i-start > lMax {
				lMax = i - start
			}
			fmt.Println(i, bMap[s[i]])

			if bMap[s[i]]+1 > start {
				start = bMap[s[i]] + 1
			}
			bMap[s[i]] = i
		}
	}

	if len(s)-start > lMax {
		lMax = len(s) - start
	}

	return lMax
}

func Test1(t *testing.T) {
	//t.Log(lengthOfLongestSubstring("abcabcbb"))
	//t.Log(lengthOfLongestSubstring("bbbbb"))
	//t.Log(lengthOfLongestSubstring("pwwkew"))
	//t.Log(lengthOfLongestSubstring(""))
	//t.Log(lengthOfLongestSubstring(" "))
	//t.Log(lengthOfLongestSubstring("aab"))
	//t.Log(lengthOfLongestSubstring("dvdf"))
	t.Log(lengthOfLongestSubstring("abba"))
}
