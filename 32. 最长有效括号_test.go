package leetcode

import "strings"

// 不会

func longestValidParentheses(s string) int {
	for {
		sTemp := s
		if strings.ReplaceAll(sTemp, "()", "") == "" {
			return len(s)
		}

		//left := strings.Count(s, "(")
		//right := strings.Count(s, ")")
		//if left > right
	}
}
