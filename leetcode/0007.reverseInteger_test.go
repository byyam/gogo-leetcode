package leetcode

import "testing"

func Test_reverseInteger(t *testing.T) {
	if rst := reverseInteger(1534236469); rst != 0 {
		t.Fatalf("failed")
	}
	t.Log("ok!")
}
