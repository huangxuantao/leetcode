package xiaohao

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func missingNumber2(nums []int) int {
	var s int
	for _, num := range nums {
		s += num
	}

	return len(nums)*(len(nums)+1)/2 - s
}

func missingNumber3(nums []int) int {
	var r int
	for i := 0; i < len(nums); i++ {
		r ^= i ^ nums[i]
	}
	return r ^ len(nums)
}
