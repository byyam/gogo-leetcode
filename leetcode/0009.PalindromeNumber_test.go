package leetcode

import "testing"

func Test_isPalindromeNumber(t *testing.T) {
	if rst := isPalindromeNumber(-10); rst {
		t.Fatalf("%v", rst)
	}
	t.Log("ok!")
}
