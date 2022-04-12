package xiaohao

func isPowerOfTwo(n int) bool {
	s := 1
	for s <= n {
		if n == s {
			return true
		}
		s = s * 2
	}
	return false
}

func isPowerOfTwo2(n int) bool {
	return n > 0 && n&(n-1) == 0
}
