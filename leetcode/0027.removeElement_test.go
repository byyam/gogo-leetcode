package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	if rst := removeElement(nums, 3); rst != 2 {
		t.Fatalf("removeElement failed:%v", rst)
	}
	t.Logf("pass:%v", nums)
}
