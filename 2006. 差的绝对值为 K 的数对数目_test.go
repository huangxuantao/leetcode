package leetcode

func countKDifference(nums []int, k int) int {
	l := len(nums)

	var c int
	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				c++
			}
		}
	}

	return c
}
