package leetcode

import "testing"

func Test_canMeasureWater(t *testing.T) {
	if rst := canMeasureWater(0, 0, 0); rst != true {
		t.Fatalf("canMeasureWater failed, rst:%v", rst)
	}
	t.Logf("canMeasureWater pass")
}
