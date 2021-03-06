package xiaohao

import (
	"testing"
)

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}

	return l
}

func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[:i+1]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func Test_removeElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	t.Log(removeElement(nums, 2))
}
