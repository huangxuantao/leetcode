package xiaohao

func minSteps(n int) int {
	var s int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			s += i
			n = n / i
		}
	}
	return s
}
