package leetcode

import "testing"

var rec1 = []int{7, 8, 13, 15}
var rec2 = []int{10, 8, 12, 20}

func Test_isRectangleOverlap(t *testing.T) {
	if rst := isRectangleOverlap(rec1, rec2); rst != true {
		t.Fatalf("isRectangleOverlap failed, rst:%v", rst)
	}
	t.Log("isRectangleOverlap pass")
}
