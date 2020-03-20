package others

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}

	sort.Ints(arr)
	return arr[0:k]
}
