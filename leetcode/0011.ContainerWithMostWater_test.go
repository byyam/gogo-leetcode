package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	if rst := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); rst != 49 {
		t.Fatalf("maxArea failed, rst:%d", rst)
	}
	t.Logf("maxArea pass")
}
