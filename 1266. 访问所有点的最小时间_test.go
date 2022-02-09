package leetcode

func minTimeToVisitAllPoints(points [][]int) int {
	var t int
	for i := 0; i < len(points)-1; i++ {
		x := points[i+1][0] - points[i][0]
		y := points[i+1][1] - points[i][1]

		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}

		if x > y {
			t += x
		} else {
			t += y
		}
	}
	return t
}
