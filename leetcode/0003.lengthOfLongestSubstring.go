package leetcode

import "fmt"

func lengthOfLongestSubstring(s string) int {
	strMap := make(map[int32]int)
	max := 0
	for index, value := range s {
		existIndex, ok := strMap[value]
		if ok {
			fmt.Printf("exist %d\n", value)
			for k, v := range strMap {
				if v <= existIndex {
					fmt.Printf("delete [%d]%d\n", v, k)
					delete(strMap, k)
				}
			}
		}
		fmt.Printf("add [%d]%d\n", index, value)
		strMap[value] = index
		if len(strMap) > max {
			fmt.Printf("update max=%d\n", len(strMap))
			max = len(strMap)
		}
	}
	return max
}
