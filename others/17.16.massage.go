package others

import "log"

func massage(nums []int) int {
	// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	log.Printf("nums:%+v", nums)
	prevMax := 0
	currMax := 0
	for i, v := range nums {
		tmp := currMax
		if currMax < prevMax+v {
			currMax = prevMax + v
		}
		log.Printf("cur=%d|%d, pre=%d, add nums[%d]=%d", tmp, currMax, prevMax, i, v)
		prevMax = tmp
	}
	return currMax
}
