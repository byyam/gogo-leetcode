package leetcode

import "testing"

func Test_minIncrementForUnique(t *testing.T) {
	if rst := minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}); rst != 6 {
		t.Fatalf("minIncrementForUnique failed:%v", rst)
	}
	t.Logf("minIncrementForUnique pass")
}
