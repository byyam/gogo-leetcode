package leetcode

import "fmt"

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	// uint64 length is 20
	l := make([]int, 20)
	length := 0
	for i := x; i > 0; i = i / 10 {
		l[length] = i % 10
		length++
	}
	fmt.Printf("x=%d, l:%v, length=%d\n", x, l, length)
	index := length - 1
	if index <= 0 {
		return true
	}
	for j := 0; j < index; j++ {
		fmt.Printf("[%d]=%d, [%d]=%d\n", j, l[j], index-j, l[index-j])
		if l[j] != l[index-j] {
			return false
		}
	}
	fmt.Println("pass")
	return true
}
