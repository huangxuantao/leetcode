package xiaohao

import "sort"

func numMovesStones(a int, b int, c int) []int {
	nums := []int{a, b, c}
	sort.Ints(nums)

	max := nums[2] - nums[0] - 2
	var min int
	aa := nums[2] - nums[1] - 1
	bb := nums[1] - nums[0] - 1

	switch {
	case aa == 0 && bb == 0:
		min = 0
	case aa > 1 && bb > 1:
		min = 2
	default:
		min = 1
	}

	return []int{min, max}
}
