package xiaohao

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	l := len(nums)

	for i := 0; i < l; i++ {
		numsI := nums[i]
		for j := i + 1; j < l; j++ {
			numsJ := nums[j]
			if numsI+numsJ == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		if v, has := numMap[target-num]; has {
			return []int{i, v}
		}
		numMap[num] = i
	}

	return nil
}

func Test_twoSum(t *testing.T) {
	t.Log(twoSum([]int{0, 4, 3, 0}, 0))
	t.Log(twoSum([]int{-3, 4, 3, 90}, 0))
	t.Log(twoSum2([]int{3, 2, 4}, 6))
	t.Log(twoSum2([]int{3, 3}, 6))
}

func Benchmark_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{0, 4, 3, 0}, 0)
	}
}
