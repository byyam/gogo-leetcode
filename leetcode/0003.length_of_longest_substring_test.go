package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	length := lengthOfLongestSubstring("dvdf")
	if length != 3 {
		t.Fatalf("lengthOfLongestSubstring error!length=%d\n", length)
	}
	t.Log("ok!")
}
