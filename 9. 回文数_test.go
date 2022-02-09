package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	l := len(s)
	t := true
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[l-1-i] {
			t = false
			break
		}
	}

	return t
}
