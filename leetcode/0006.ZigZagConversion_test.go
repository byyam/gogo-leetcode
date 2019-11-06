package leetcode

import "testing"

func Test_ZigZagConversion(t *testing.T) {
	if rst := ZigZagConversion("PAYPALISHIRING", 3); rst != "PAHNAPLSIIGYIR" {
		t.Fatalf("return %s\n", rst)
	}
	t.Log("ok!")
}
