package leetcode

import "fmt"

var romanMapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if len(s) <= 1 {
		return romanMapping[s]
	}

	result := 0
	var id int
	for id = 0; id < len(s)-1; id++ {
		thisOne := string(s[id])
		nextOne := string(s[id+1])
		// get 4 or 9, use two char
		if romanMapping[thisOne] < romanMapping[nextOne] {
			id++
			result = result + romanMapping[nextOne] - romanMapping[thisOne]
		} else {
			result = result + romanMapping[thisOne]
		}
		fmt.Printf("[%d]this:%s, next:%s, result:%d\n", id, thisOne, nextOne, result)
	}
	if id == len(s)-1 {
		lastOne := string(s[id])
		result = result + romanMapping[lastOne]
		fmt.Printf("last:%s, result:%d\n", lastOne, result)
	}

	return result
}
