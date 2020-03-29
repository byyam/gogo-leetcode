package leetcode

import "testing"

var maxDistance_test = [][]int{
	{1, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

func Test_maxDistance(t *testing.T) {
	if rst := maxDistance(maxDistance_test); rst != 4 {
		t.Fatalf("max distance failed:%d", rst)
	}
	t.Logf("max distance pass")
}
