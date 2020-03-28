package leetcode

import (
	"log"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	// 先将words按字符串长度进行排序
	maxLen := 0
	for _, v := range words {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	log.Printf("max length of words is:%d", maxLen)
	if maxLen == 0 {
		return 0
	}

	sortedWords := []string{}
	for i := 0; i <= maxLen; i++ {
		for _, v := range words {
			if len(v) == i {
				sortedWords = append(sortedWords, v)
			}
		}
	}
	log.Printf("sorted words is:%+v", sortedWords)

	// 如果word1是word2的字串, 且word1必须是在word2的结尾处，则word1的长度可以被忽略
	resultWords := []string{}
	for i := 0; i < len(sortedWords)-1; i++ {
		flag := false
		for j := i + 1; j < len(sortedWords); j++ {
			word1 := sortedWords[i]
			word2 := sortedWords[j]
			if strings.HasSuffix(word2, word1) {
				log.Printf("word1:%s, word2:%s", word1, word2)
				flag = true
			}
		}
		if !flag {
			resultWords = append(resultWords, sortedWords[i])
		}
	}
	resultWords = append(resultWords, sortedWords[len(sortedWords)-1])
	log.Printf("result words is:%+v", resultWords)

	sum := 0
	for _, v := range resultWords {
		sum += len(v) + 1
	}

	return sum
}
