package xiaohao

import (
	"sort"
	"testing"
)

func twoSum(nums []int, target int) []int {
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l; i++ {
		numsI := nums[i]
		if numsI > target {
			continue
		}

		for j := i + 1; j < l; j++ {
			numsJ := nums[j]
			if numsJ > target {
				break
			}

			if numsI+numsJ == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func Test_twoSum(t *testing.T) {
	t.Log(twoSum([]int{0, 4, 3, 0}, 0))
	t.Log(twoSum([]int{-3, 4, 3, 90}, 0))
}

func Benchmark_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{0, 4, 3, 0}, 0)
	}
}
