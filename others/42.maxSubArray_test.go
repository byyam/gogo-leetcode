package others

import "testing"

func Test_maxSubArray(t *testing.T) {
	max := maxSubArray([]int{1})
	t.Logf("max:%d", max)
}
