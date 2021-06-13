package leetcode

import (
	"sort"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aList := append(nums1, nums2...)
	sort.Ints(aList)

	if len(aList)%2 == 1 {
		return float64(aList[len(aList)/2])
	} else {
		return (float64(aList[len(aList)/2-1]) + float64(aList[len(aList)/2])) / 2
	}
}

func TestName(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 3}, []int{2}))
	t.Log(findMedianSortedArrays([]int{1, 1}, []int{1, 2}))
}
