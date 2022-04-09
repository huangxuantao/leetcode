package leetcode

func search(nums []int, target int) int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		numMap[num] = i
	}

	if i, has := numMap[target]; has {
		return i
	}
	return -1
}
