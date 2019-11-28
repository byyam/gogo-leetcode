package leetcode

func TwoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for index, value := range nums {
		anotherValue := target - value
		anotherIndex, exist := sumMap[anotherValue]
		if exist {
			return []int{index, anotherIndex}
		}
		sumMap[value] = index
	}
	return []int{}
}
