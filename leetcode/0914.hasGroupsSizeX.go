package leetcode

import (
	"log"
	"sort"
)

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	// 先把列表放在map里，key为值，value为重复次数
	repeatMap := make(map[int]int)
	for _, v := range deck {
		repeatMap[v] += 1
	}
	log.Printf("repeat map:%+v", repeatMap)

	// 再把map中value相同的值过滤掉，即去掉重复次数相同的
	countList := make([]int, 0)
	countMap := make(map[int]bool)
	for _, v := range repeatMap {
		if countMap[v] == true {
			continue
		} else {
			countMap[v] = true
			countList = append(countList, v)
		}
	}
	sort.Ints(countList)
	log.Printf("count list:%+v", countList)

	for i := 2; i <= countList[0]; i++ {
		flag := true
		for _, v := range countList {
			if v%i != 0 {
				flag = false
			}
		}
		log.Printf("%d is %v", i, flag)
		if flag {
			return true
		}
	}

	return false
}
