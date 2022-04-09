package leetcode

import (
	"testing"
)

func searchRange(nums []int, target int) []int {
	numsOrigin := nums
	l := len(nums)
	lOrigin := l

	if l == 0 {
		return []int{-1, -1}
	}

	iStart := 0
	iEnd := l - 1

	for iStart < iEnd-1 {
		i := iStart + (iEnd-iStart)/2
		// fmt.Println(iStart, iEnd, i)

		switch {
		case nums[i] > target:
			iEnd = i

		case nums[i] < target:
			iStart = i

		default:
			iStart, iEnd = i, i
			for {
				if iStart > 0 && numsOrigin[iStart-1] == target {
					iStart--
					continue
				}
				break
			}

			for {
				if iEnd < lOrigin-1 && numsOrigin[iEnd+1] == target {
					iEnd++
					continue
				}
				break
			}

			return []int{iStart, iEnd}
		}
	}

	if nums[iStart] == target && nums[iEnd] == target {
		return []int{iStart, iEnd}
	}

	if nums[iStart] == target {
		return []int{iStart, iStart}
	}
	if nums[iEnd] == target {
		return []int{iEnd, iEnd}
	}
	return []int{-1, -1}

}

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	t.Log(searchRange(nums, 8))

	nums = []int{5, 7, 7, 8, 8, 10}
	t.Log(searchRange(nums, 6))

	nums = []int{1, 2, 3}
	t.Log(searchRange(nums, 3))

	nums = []int{1, 3}
	t.Log(searchRange(nums, 1))

	nums = []int{1, 2, 3}
	t.Log(searchRange(nums, 3))

	nums = []int{0, 1, 2, 3, 4, 4, 4}
	t.Log(searchRange(nums, 2))
}
