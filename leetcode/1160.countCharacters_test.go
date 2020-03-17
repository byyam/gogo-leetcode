package leetcode

import "testing"

var word = []string{"cat", "bt", "hat", "tree"}
var chars = "atach"
var result = 6

func Test_countCharacters(t *testing.T) {
	if rst := countCharacters(word, chars); rst != result {
		t.Fatalf("countCharacters failed:%d", rst)
	}
	t.Logf("countCharacters pass")
}
