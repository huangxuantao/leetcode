package leetcode

import (
	"strings"
	"testing"
)

func intToRoman(num int) string {
	g := make(map[int]string)
	s := make(map[int]string)
	b := make(map[int]string)

	g[1] = "I"
	g[2] = "II"
	g[3] = "III"
	g[4] = "IV"
	g[5] = "V"
	g[6] = "VI"
	g[7] = "VII"
	g[8] = "VIII"
	g[9] = "IX"
	s[1] = "X"
	s[2] = "XX"
	s[3] = "XXX"
	s[4] = "XL"
	s[5] = "L"
	s[6] = "LX"
	s[7] = "LXX"
	s[8] = "LXXX"
	s[9] = "XC"
	b[1] = "C"
	b[2] = "CC"
	b[3] = "CCC"
	b[4] = "CD"
	b[5] = "D"
	b[6] = "DC"
	b[7] = "DCC"
	b[8] = "DCCC"
	b[9] = "CM"

	builder := strings.Builder{}
	switch num / 1000 {
	case 1:
		builder.WriteString("M")
	case 2:
		builder.WriteString("MM")
	case 3:
		builder.WriteString("MMM")
	default:
	}

	builder.WriteString(b[num%1000/100])
	builder.WriteString(s[num%100/10])
	builder.WriteString(g[num%10])

	return builder.String()

}

func Test_intToRoman(t *testing.T) {
	t.Log(intToRoman(58))
	t.Log(intToRoman(1994))
}
