package leetcode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	combine := make([]int, len(nums1)+len(nums2))
	fmt.Printf("len(combine)=%d,num1:%v,num2:%v\n", len(combine), nums1, nums2)
	if len(combine) == 0 {
		return 0
	}
	x, y := 0, 0
	for z := 0; z < len(combine); z++ {
		if (x == len(nums1)) && (y == len(nums2)) {
			break
		} else if x == len(nums1) {
			combine[z] = nums2[y]
			y++
		} else if y == len(nums2) {
			combine[z] = nums1[x]
			x++
		} else if nums1[x] > nums2[y] {
			combine[z] = nums2[y]
			y++
		} else {
			combine[z] = nums1[x]
			x++
		}
		fmt.Printf("x=%d,y=%d,z=%d,combine:%v\n", x, y, z, combine)
	}
	var result float64
	if len(combine)%2 == 1 {
		result = float64(combine[(len(combine)-1)/2])
	} else {
		result = (float64(combine[len(combine)/2-1]) + float64(combine[len(combine)/2])) / 2
	}
	fmt.Printf("result=%f\n", result)
	return result
}
