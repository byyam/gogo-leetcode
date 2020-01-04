package leetcode

import (
	"fmt"
)

func maxArea(height []int) int {
	length := len(height)
	if length < 2 {
		return 0
	}
	if length == 2 {
		high := height[0]
		if height[0] > height[1] {
			high = height[1]
		}
		return high
	}
	fmt.Printf("length=%d\n", length)
	max := 0
	for start := 0; start < length-1; start++ {
		for end := start + 1; end < length; end++ {
			wide := end - start
			high := height[start]
			if height[start] > height[end] {
				high = height[end]
			}
			area := high * wide
			if max < area {
				max = area
			}
			fmt.Printf("start:%d,end:%d,high:%d,wide:%d,area:%d\n", start, end, high, wide, area)
		}
	}
	return max
}
