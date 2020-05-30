package others

/*
dp(n) = dp(n-1) + num(n)  [dp(n-1) >= 0]
dp(n) = num(n)  [dp(n-1) < 0]
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	dp := 0
	for _, n := range nums {
		if dp < 0 {
			dp = 0
		}
		dp += n
		if dp > max {
			max = dp
		}
	}
	return max
}
