package xiaohao

import "math"

func integerBreak(n int) int {
	if n== 2 {
		return 1
	}
	if n==3 {
		return 2
	}

	if n%3 == 1 {
		return int(math.Pow(3, float64(n/3)-1) * 2 * 2)
	} else if n%3 == 2 {
		return int(math.Pow(3, float64(n/3)) * 2)
	} else {
		return int(math.Pow(3, float64(n/3)))
	}
}
