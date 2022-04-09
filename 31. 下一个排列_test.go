package leetcode

// 不会

import (
	"sort"
	"testing"
)

func nextPermutation(nums []int) {
	l := len(nums)
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			t := nums[i]
			for j := i; j < l-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[l-1] = t

			//fmt.Println(nums)
			return
		}
	}

	sort.Ints(nums)
}
func Test_nextPermutation(t *testing.T) {
	nextPermutation([]int{1, 2, 3, 4, 5})
	nextPermutation([]int{1, 3, 2})
}
