package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	retMap := make(map[string][]int)

	sort.Ints(nums)
	//fmt.Println(nums)

	l := len(nums)
	for i := 0; i < l-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}

		if i > 0 && n1 == nums[i-1] {
			continue
		}

		left := i + 1
		right := l - 1

		for left < right {
			n2 := nums[left]
			n3 := nums[right]

			sum := n1 + n2 + n3

			if sum == 0 {
				retMap[fmt.Sprintf("%d%d%d", n1, n2, n3)] = []int{n1, n2, n3}
				for left < right && n2 == nums[left] {
					left++
				}

				for left < right && n3 == nums[right] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	for _, i := range retMap {
		ret = append(ret, i)
	}

	return ret
}

func Test_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))

	nums = []int{0, 0, 0}
	t.Log(threeSum(nums))

	nums = []int{-2, 0, 1, 1, 2}
	t.Log(threeSum(nums))

	nums = []int{1, 2, -2, -1}
	t.Log(threeSum(nums))
}
