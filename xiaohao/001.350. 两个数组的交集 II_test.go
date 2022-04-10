package xiaohao

import "testing"

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)

	for _, i := range nums1 {
		nums1Map[i]++
	}

	var nums3 []int

	for _, i := range nums2 {
		if nums1Map[i] > 0 {
			nums1Map[i]--
			nums3 = append(nums3, i)
		}
	}

	return nums3
}

func Test_intersect(t *testing.T) {

}
