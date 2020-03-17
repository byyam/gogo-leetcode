package leetcode

import "log"

func countCharacters(words []string, chars string) int {
	sum := 0

	for _, word := range words {
		dict := make(map[int32]int)
		for _, v := range chars {
			dict[v] = dict[v] + 1
		}
		set := true
		for _, w := range word {
			v, exist := dict[w]
			if !exist || v == 0 {
				set = false
				log.Printf("value:%c in word:%s not exist:%v, v=%d", w, word, exist, v)
				break
			}
			dict[w] = v - 1
		}
		if set {
			sum = sum + len(word)
			log.Printf("word:%s in dict, sum=%d", word, sum)
		}
	}

	return sum
}
