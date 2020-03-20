package others

import "testing"

func Test_getLeastNumbers(t *testing.T) {
	rst := getLeastNumbers([]int{0, 1, 2, 1}, 1)
	t.Logf("get least numbers:%v", rst)
}
