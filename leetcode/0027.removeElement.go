package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	rst := 0
	change := len(nums) - 1
	for rst < change {
		fmt.Printf("value=%d,val=%d,index=%d,change=%d,nums:%v\n", nums[rst], val, rst, change, nums)
		if nums[rst] != val {
			fmt.Printf("rst++,index=%d\n", rst)
			rst++
		} else {
			fmt.Printf("change--,index=%d\n", rst)
			nums[rst], nums[change] = nums[change], nums[rst]
			change--
		}
	}
	if nums[rst] != val {
		rst++
	}
	return rst
}
