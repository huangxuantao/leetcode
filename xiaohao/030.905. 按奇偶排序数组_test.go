package xiaohao

func sortArrayByParity(nums []int) []int {
	var a, b []int
	for _, num := range nums {
		if num%2 == 0 {
			a = append(a, num)
		} else {
			b = append(b, num)
		}
	}

	return append(a, b...)
}

func sortArrayByParity2(nums []int) []int {
	var j int
	for i, num := range nums {
		if num%2==0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return nums
}
