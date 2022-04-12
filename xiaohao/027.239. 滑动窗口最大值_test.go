package xiaohao

func maxSlidingWindow(nums []int, k int) []int {
	var q []int
	var r []int

	for i, num := range nums {
		for i > 0 && len(q) > 0 && num > q[len(q)-1] {
			q = q[:len(q)-1]
		}

		q = append(q, num)
		if i >= k && nums[i-k] == q[0] {
			q = q[1:]
		}

		if i >= k-1 {
			r = append(r, q[0])
		}
	}

	return r
}
