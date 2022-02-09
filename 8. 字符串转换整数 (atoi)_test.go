package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func myAtoi(s string) int {
	builder := strings.Builder{}

	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	if s[0] == '-' {
		for i := 1; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			if err == nil {
				builder.WriteByte(s[i])
			} else {
				break
			}
		}

		if builder.String() == "" {
			return 0
		}

		a, err := strconv.Atoi(builder.String())
		if err != nil {
			if strings.Contains(err.Error(), "out of range") {
				return -math.MaxInt32 - 1
			}
			return 0
		}

		if a > math.MaxInt32 {
			return -math.MaxInt32 - 1
		}

		return -a

	} else if s[0] == '+' {
		for i := 1; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			if err == nil {
				builder.WriteByte(s[i])
			} else {
				break
			}
		}

		if builder.String() == "" {
			return 0
		}

		a, err := strconv.Atoi(builder.String())
		if err != nil {
			if strings.Contains(err.Error(), "out of range") {
				return math.MaxInt32
			}
			return 0
		}

		if a > math.MaxInt32 {
			return math.MaxInt32
		}

		return a

	} else {
		for i := 0; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			if err == nil {
				builder.WriteByte(s[i])
			} else {
				break
			}
		}

		fmt.Println(builder.String())
		if builder.String() == "" {
			return 0
		}

		a, err := strconv.Atoi(builder.String())
		if err != nil {
			if strings.Contains(err.Error(), "out of range") {
				return math.MaxInt32
			}
			return 0
		}

		if a > math.MaxInt32 {
			return math.MaxInt32
		}

		return a
	}

}

func Test_myAtoi(t *testing.T) {
	t.Log(myAtoi("42"))
	t.Log(myAtoi("   -42"))
	t.Log(myAtoi("20000000000000000000"))
}
