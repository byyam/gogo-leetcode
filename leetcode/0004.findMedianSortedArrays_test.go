package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	if rst := findMedianSortedArrays([]int{1, 2, 3}, []int{4}); rst != 2.5 {
		t.Fatalf("findMedianSortedArrays failed!\n")
	}
	t.Log("ok!")
}
