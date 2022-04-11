package xiaohao

func minPathSum(grid [][]int) int {
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

	m := len(grid)
	n := len(grid[0])

	dp := make(map[p]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[p{i, i}] = grid[0][0]
			} else if i == 0 {
				dp[p{i, j}] = dp[p{i, j - 1}] + grid[i][j]
			} else if j == 0 {
				dp[p{i, j}] = dp[p{i - 1, j}] + grid[i][j]
			} else {
				dp[p{i, j}] = minFunc(dp[p{i, j - 1}]+grid[i][j], dp[p{i - 1, j}]+grid[i][j])
			}
		}
	}

	return dp[p{m - 1, n - 1}]
}
