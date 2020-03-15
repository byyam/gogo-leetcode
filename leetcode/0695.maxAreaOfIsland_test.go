package leetcode

import "testing"

var grid = [][]int{
	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

func Test_maxAreaOfIsland(t *testing.T) {
	if rst := maxAreaOfIsland(grid); rst != 6 {
		t.Fatalf("max area of island failed:%v", rst)
	}
	t.Logf("max area of island pass")
}
