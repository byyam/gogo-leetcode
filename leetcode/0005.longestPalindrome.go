package leetcode

import "fmt"

func longestPalindrome(s string) string {
	var dp [1000][1000]bool
	length := len(s)
	if length <= 1 {
		return s
	}
	x, y := 0, 0
	fmt.Printf("%s\n", s)
	for right := 1; right < length; right++ {
		// range: [0, right]
		for left := 0; left < right; left++ {
			// range: [left, right]
			if s[left] == s[right] && (dp[left+1][right-1] || (right-1) <= (left+1)) {
				dp[left][right] = true
				if (right - left) > (y - x) {
					y = right
					x = left
				}
			}
			fmt.Printf("dp[%d][%d]=%v, %s, subStr:%s\n", left, right, dp[left][right], s[left:right+1], s[x:y+1])
		}
	}
	subStr := s[x : y+1]
	fmt.Printf("s[%d:%d],substr=%s\n", x, y+1, subStr)
	return subStr
}
