package toolkit

import "fmt"

func PrintGrid(grid [][]int) {
	for _, list := range grid {
		fmt.Printf("%v\n", list)
	}
}

func CompareGrid(grid1, grid2 [][]int) bool {
	if len(grid1) != len(grid2) {
		return false
	}
	for index, array1 := range grid1 {
		array2 := grid2[index]
		if !CompareArray(array1, array2) {
			return false
		}
	}
	return true
}

func CompareArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for index, value1 := range arr1 {
		value2 := arr2[index]
		if value1 != value2 {
			return false
		}
	}
	return true
}
