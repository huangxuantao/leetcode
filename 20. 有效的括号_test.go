package leetcode

import (
	"strings"
	"testing"
)

func isValid(s string) bool {
	a := true
	b := true
	c := true

	for a || b || c {
		if strings.Index(s, "[]") == -1 {
			a = false
		} else {
			a = true
			s = strings.ReplaceAll(s, "[]", "")
		}

		if strings.Index(s, "{}") == -1 {
			b = false
		} else {
			b = true
			s = strings.ReplaceAll(s, "{}", "")
		}

		if strings.Index(s, "()") == -1 {
			c = false
		} else {
			c = true
			s = strings.ReplaceAll(s, "()", "")
		}
	}

	if s == "" {
		return true
	}
	return false
}

func Test_isValid(t *testing.T) {
	s := "()"
	t.Log(isValid(s))
}
