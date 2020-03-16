package others

import "testing"

func Test_compressString(t *testing.T) {
	if rst := compressString("aabcccccaaa"); rst != "a2b1c5a3" {
		t.Fatalf("compress string failed:%s", rst)
	}
	t.Logf("compress string pass")
}
