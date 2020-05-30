package others

import "testing"

func Test_missingNumber(t *testing.T) {
	rst := missingNumber([]int{0})
	t.Logf("missing number is:%d", rst)
}
