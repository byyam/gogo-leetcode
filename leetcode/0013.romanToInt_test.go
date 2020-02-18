package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
	if rst := romanToInt("MCMXCIV"); rst != 1994 {
		t.Fatalf("romonToInt failed:%d", rst)
	}
	t.Logf("romonToInt pass")
}
