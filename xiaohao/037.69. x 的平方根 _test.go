package xiaohao

import (
	"math"
	"testing"
)

func mySqrt(x int) int {
	min := 0
	max := math.MaxInt32

	var c int
	for {
		if max-min == 1 {
			return min
		}

		//fmt.Println(min, max, c)
		c = (min + max) / 2
		ca := c * c
		cb := (c + 1) * (c + 1)

		if ca == x {
			return c
		}

		if cb == x {
			return c + 1
		}

		if ca < x {
			min = c
		} else if cb > x {
			max = c
		} else {
			return c
		}
	}
}

func Test_mySqrt(t *testing.T) {
	t.Log(mySqrt(2))
	t.Log(mySqrt(3))
	t.Log(mySqrt(4))
	t.Log(mySqrt(5))
}
