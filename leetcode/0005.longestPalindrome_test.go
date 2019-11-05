package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	if rst := longestPalindrome("babad"); rst != "bab" && rst != "aba" {
		t.Fatalf("wrong result:%s\n", rst)
	}
	t.Log("pass")
}
