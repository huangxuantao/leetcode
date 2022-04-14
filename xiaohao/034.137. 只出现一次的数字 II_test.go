package xiaohao

import "sort"

func singleNumber11(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i+=3 {
		if nums[i]!=nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}