package leetcode

import (
	"log"
	"sort"
)

func threeSum(nums []int) [][]int {
	var answer [][]int
	if len(nums) < 3 { // 异常情况，长度小于3
		return answer
	}
	sort.Ints(nums) // 从小到大排序
	log.Printf("sort nums:%+v", nums)
	for index, elem := range nums { // elem 为目标数，另外两个左右指针是移动的，区间为elem到最大的数
		if elem > 0 { // 目标数从左边开始，如果目标数已经大于0，后面都比elem大，就不会有合为0的情况了，直接退出循环
			return answer
		}

		if index > 0 {
			if elem == nums[index-1] { // 目标数如果相同，会导致重复结果
				continue
			}
		}

		L := index + 1
		R := len(nums) - 1
		for L < R { // 左右区间 [elem+1, last]
			if elem+nums[L]+nums[R] == 0 { // 满足条件还要继续找，可能会有多组满足条件的值
				answer = append(answer, append([]int{}, elem, nums[L], nums[R]))
				for L < R && nums[L] == nums[L+1] { // 跳过重复值
					L += 1
				}
				for L < R && nums[R] == nums[R-1] { // 跳过重复值
					R -= 1
				}
				L += 1
				R -= 1
			} else if elem+nums[L]+nums[R] > 0 { // 若结果偏大，则移动右边使结果缩小
				R -= 1
			} else if elem+nums[L]+nums[R] < 0 { // 反之
				L += 1
			}
		}
	}
	return answer
}
