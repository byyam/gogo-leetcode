package leetcode

import "testing"

func Test_knightProbability(t *testing.T) {
	if rst := knightProbability(3, 2, 0, 0); rst != 0.0625 {
		t.Fatalf("rst=%f", rst)
	}
	t.Log("ok!")
}
