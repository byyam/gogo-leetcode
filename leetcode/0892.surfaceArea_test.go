package leetcode

import "testing"

var surfaceAreaGrid = [][]int{
	{1, 0},
	{0, 2},
}

func Test_surfaceArea(t *testing.T) {
	if rst := surfaceArea(surfaceAreaGrid); rst != 16 {
		t.Fatalf("surface area failed:%d", rst)
	}
	t.Logf("surface area pass")
}
