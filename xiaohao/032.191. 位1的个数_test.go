package xiaohao

import (
	"fmt"
	"strings"
	"testing"
)

func hammingWeight(num uint32) int {
	s := fmt.Sprintf("%b", num)
	// fmt.Println(s)
	return strings.Count(s, "1")
}

func hammingWeight2(num uint32) int {
	var ret int
	for num != 0 {
		if num&0b1 == 1 {
			ret++
		}
		num = num >> 1
	}
	return ret
}

func Test_hammingWeight(t *testing.T) {
	t.Log(hammingWeight2(8))
	t.Log(hammingWeight2(7))
}
