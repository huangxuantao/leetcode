package xiaohao

func plusOne(digits []int) []int {
	l := len(digits)
	k := 0
	for {
		if k == l {
			return append([]int{1}, digits...)
		}

		c := digits[l-1-k]
		c++

		if c == 10 {
			digits[l-1-k] = 0
			k++
		} else {
			digits[l-1-k] = c
			break
		}
	}

	return digits
}
