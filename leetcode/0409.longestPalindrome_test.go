package leetcode

import "testing"

func Test_longestPalindromeLength(t *testing.T) {
	if rst := longestPalindromeLength("abccccdd"); rst != 7 {
		t.Fatalf("longestPalindromeLength failed, rst=%d", rst)
	}
	t.Logf("longestPalindromeLength pass")
}
