package leetcode

import (
	"log"
	"sort"
)

func minIncrementForUnique(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	sum := 0
	// 排序
	sort.Ints(A)
	log.Printf("sort array:%v", A)

	for i, _ := range A {
		if i == 0 {
			continue
		}
		for A[i] <= A[i-1] {
			A[i]++
			sum++
			log.Printf("A:%v", A)
		}
	}

	return sum
}
