package leetcode

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func reverse(x int) int {
	f := false
	if x < 0 {
		f = true
	}

	builder := strings.Builder{}

	var xx string
	if f {
		xx = strconv.Itoa(-x)
	} else {
		xx = strconv.Itoa(x)
	}

	//fmt.Println(xx)
	for i := len(xx) - 1; i >= 0; i-- {
		builder.WriteByte(xx[i])
	}

	//fmt.Println(builder.String())
	y, _ := strconv.Atoi(builder.String())

	if y > math.MaxInt32 {
		return 0
	}

	if f {
		return -y
	} else {
		return y
	}
}

func Test_reverse(t *testing.T) {
	t.Log(reverse(123))
	t.Log(reverse(-123))
	t.Log(reverse(1534236469))
}

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123)
	}
}
