package xiaohao

func maxSubArray(nums []int) int {
	dp := make(map[int]int)
	dp[0] = nums[0]

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		var v int
		if nums[i]+dp[i-1] > nums[i] {
			v = nums[i] + dp[i-1]
		} else {
			v = nums[i]

		}
		dp[i] = v
		if v > max {
			max = v
		}
	}

	return max
}
