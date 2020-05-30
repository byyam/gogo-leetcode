package others

/*
1 + 2 + 3 = 3 * (1+3)/2
sum(0~n-1) = n*(n-1 + 0) / 2
*/
func missingNumber(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	length := len(nums) + 1
	sum := length * (length - 1) / 2
	for _, n := range nums {
		sum = sum - n
	}
	return sum
}
