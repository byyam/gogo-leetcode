package leetcode

import (
	toolkit "github.com/byyam/go-leetcode-toolkit"
	"testing"
)

var input = []int{-1, 0, 1, 2, -1, -4}

func Test_threeSum(t *testing.T) {
	rst := threeSum(input)

	toolkit.PrintGrid(rst)
}
