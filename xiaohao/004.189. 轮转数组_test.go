package xiaohao

func rotate(nums []int, k int) {
	var a []int
	l := len(nums)
	k = k % l

	a = append(a, nums[l-k:]...)
	a = append(a, nums[:l-k]...)
	for i := 0; i < l; i++ {
		nums[i] = a[i]
	}
}

func rotate2(nums []int, k int) {
	r := func(nums []int) {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]

		}
	}

	l := len(nums)
	k = k % l

	r(nums)
	r(nums[k:])
	r(nums[:k])
}
