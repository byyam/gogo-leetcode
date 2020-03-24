package others

import "testing"

func Test_massage(t *testing.T) {
	if rst := massage([]int{1, 2, 3, 1}); rst != 4 {
		t.Fatalf("massage failed:%v", rst)
	}
	t.Logf("massage pass")
}
