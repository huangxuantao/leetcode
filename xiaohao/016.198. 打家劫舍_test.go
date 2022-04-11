package xiaohao

func rob(nums []int) int {
	maxFunc := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	if len(nums) < 2 {
		return nums[0]
	}

	max := nums[0]
	dp := make(map[int]int)
	dp[0] = nums[0]
	dp[1] = maxFunc(dp[0], nums[1])
	if max < dp[1] {
		max = dp[1]
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = maxFunc(dp[i-2]+nums[i], dp[i-1])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
