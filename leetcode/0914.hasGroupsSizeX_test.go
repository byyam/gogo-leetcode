package leetcode

import "testing"

func Test_hasGroupsSizeX(t *testing.T) {
	if rst := hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2, 3, 3}); rst != true {
		t.Fatalf("hasGroupsSizeX failed:%v", rst)
	}
	t.Logf("hasGroupsSizeX pass")
}
