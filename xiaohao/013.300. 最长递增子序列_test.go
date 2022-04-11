package xiaohao

import (
	"fmt"
	"testing"
)

func lengthOfLIS(nums []int) int {
	maxFunc := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[i] = num
	}

	dp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxFunc(dp[j]+1, dp[i])
			}
		}
	}

	max := 1
	for _, i := range dp {
		if max < i {
			max = i
		}
	}

	fmt.Println(dp)
	return max
}

func Test_lengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{4,10,4,3,8,9}
	t.Log(lengthOfLIS(nums))
}
