package xiaohao

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)

	var sub int
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			continue
		}

		if nums[i+1] == nums[i] {
			return false
		}

		sub += nums[i+1] - nums[i]
	}

	return sub < 5

}
