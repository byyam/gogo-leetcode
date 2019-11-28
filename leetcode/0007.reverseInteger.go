package leetcode

import (
	"fmt"
	"math"
)

func reverseInteger(x int) int {
	negative := false
	fmt.Printf("x=%d\n", x)
	if x < 0 {
		negative = true
		x = x * (-1)
		fmt.Printf("negative:%v, x=%d\n", negative, x)
	}
	remainder := x
	result := 0
	fmt.Printf("before, remainder=%d, result=%d\n", remainder, result)
	for remainder > 0 {
		result = remainder%10 + result*10
		remainder = remainder / 10
		fmt.Printf("after,  remainder=%d, result=%d\n", remainder, result)
	}
	if result > math.MaxInt32 {
		result = 0
	}
	if negative {
		result = result * (-1)
	}
	fmt.Printf("result=%d\n", result)
	return result
}
