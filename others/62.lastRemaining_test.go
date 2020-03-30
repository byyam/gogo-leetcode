package others

import (
	"testing"
)

func Test_lastRemaining(t *testing.T) {
	if rst := lastRemaining(5, 3); rst != 3 {
		t.Fatalf("last remaining failed:%d", rst)
	}
	t.Logf("last remaining pass")
}
