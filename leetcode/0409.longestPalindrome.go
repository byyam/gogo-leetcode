package leetcode

func longestPalindromeLength(s string) int {
	// 相同字符，则可以构成回文
	// 把字符放在哈希里，统计出现次数
	dict := make(map[int32]int)
	for _, v := range s {
		dict[v] += 1
	}

	sum := 0
	for _, v := range dict {
		sum += v / 2 * 2 // 只计算偶数部分
	}
	// 统计完对称字符，中间可再放入任意一个字符
	if sum != len(s) {
		sum++
	}

	return sum
}
