package others

import (
	"fmt"
	"strings"
)

func compressString(input string) string {
	if len(input) == 0 {
		return ""
	}
	output := ""
	cache := int32(input[0]) // 压缩字符初始取第一个字符
	count := 0               // 压缩字符相同个数初始化为0
	for _, value := range input {
		if cache == value {
			count++
		} else {
			output = strings.Join([]string{output, fmt.Sprintf("%c%d", cache, count)}, "")
			cache = value
			count = 1
		}
	}
	output = strings.Join([]string{output, fmt.Sprintf("%c%d", cache, count)}, "")

	if len(output) >= len(input) {
		return input
	}
	return output
}
