package xiaohao

import (
	"sort"
	"testing"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func singleNumber2(nums []int) int {
	var r int
	for _, num := range nums {
		r ^= num
	}
	return r
}

func Test_singleNumber(t *testing.T) {
	t.Log(singleNumber([]int{1, 2, 2, 3, 3}))
	t.Log(singleNumber2([]int{1, 2, 2, 3, 3}))
}
