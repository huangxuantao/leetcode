package xiaohao

import (
	"testing"
)

func maxProfit(prices []int) int {
	s := 0
	for i := 0; i < len(prices)-1; i++ {
		d := prices[i+1] - prices[i]
		if d >= 0 {
			s += d
		}

	}

	return s
}

func Test_maxProfit(t *testing.T) {
	p := []int{7, 1, 5, 3, 6, 4}
	t.Log(maxProfit(p))

	p = []int{1, 2, 3, 4, 5}
	t.Log(maxProfit(p))
}
