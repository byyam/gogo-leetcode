package leetcode

import "fmt"

func ZigZagConversion(s string, numRows int) string {
	length := len(s)
	if length <= numRows || numRows <= 1 {
		fmt.Printf("num=%d\n", numRows)
		return s
	}
	str := ""
	loop := numRows + numRows - 2
	fmt.Printf("s:%s,length=%d,loop=%d\n", s, length, loop)
	for l := 0; l < numRows; l++ {
		for i := l; i < length; {
			str = fmt.Sprintf("%s%c", str, s[i])
			fmt.Printf("[%d]add %c, i=%d,str=%s\n", l, s[i], i, str)
			if l != 0 && l != numRows-1 {
				zip := i + loop - 2*l
				fmt.Printf("zip=%d\n", zip)
				if zip < length {
					str = fmt.Sprintf("%s%c", str, s[zip])
					fmt.Printf("[%d]add %c, i=%d,str=%s\n", l, s[zip], zip, str)
				}
			}
			i = i + loop
		}
	}
	return str
}
