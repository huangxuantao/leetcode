package xiaohao

import (
	"testing"
)

func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := 0
	minSpeed := 0
	for _, pile := range piles {
		if max < pile {
			max = pile
		}
	}

	for {
		minSpeed = (min + max) / 2
		hh := 0
		for _, pile := range piles {
			hh += pile / minSpeed
			if pile%minSpeed != 0 {
				hh++
			}
		}

		if hh > h {
			if min == minSpeed {
				min++
				minSpeed++
			} else {
				min = minSpeed
			}
			//fmt.Println("min", min)

		} else if hh < h {
			if max == minSpeed {
				max--
				minSpeed--
			} else {
				max = minSpeed
			}
			//fmt.Println("max", max)

		} else {
			max = minSpeed
		}

		if max <= min {
			break
		}
	}

	return minSpeed
}

func Test_minEatingSpeed(t *testing.T) {
	t.Log(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	t.Log(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	t.Log(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	t.Log(minEatingSpeed([]int{1, 1, 1, 999999999}, 10))
}
