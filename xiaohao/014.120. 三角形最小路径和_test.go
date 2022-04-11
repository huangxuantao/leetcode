package xiaohao

func minimumTotal(triangle [][]int) int {
	minFunc := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	type p struct {
		I int
		J int
	}

	l := len(triangle)
	dp := make(map[p]int)
	dp[p{0, 0}] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for ii, num := range triangle[i] {
			if ii == 0 {
				dp[p{i, ii}] = dp[p{i - 1, ii}] + num
			} else if ii == len(triangle[i])-1 {
				dp[p{i, ii}] = dp[p{i - 1, ii - 1}] + num
			} else {
				dp[p{i, ii}] = minFunc(dp[p{i - 1, ii - 1}]+num, dp[p{i - 1, ii}]+num)
			}
		}
	}

	min := dp[p{l-1, 0}]
	for k, v := range dp {
		if k.I != l-1 {
			continue
		}
		if v < min {
			min = v
		}
	}

	return min
}
